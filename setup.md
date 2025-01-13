# OpenComply AWS Integration

This repository provides resources and instructions for integrating OpenComply with AWS, supporting both single-account and multi-account (AWS Organization) setups.

## Table of Contents

1. [Overview](#overview)
2. [Prerequisites Checklist](#prerequisites-checklist)
3. [Integration Methods](#integration-methods)
    - [1. AWS Single Account (Console)](#1-aws-single-account-console)
    - [2. AWS Single Account (CLI)](#2-aws-single-account-cli)
    - [3. AWS Multi-Account / Organization (CLI)](#3-aws-multi-account--organization-cli)
4. [Security Considerations](#security-considerations)
5. [Troubleshooting](#troubleshooting)
6. [Support](#support)

## Overview

OpenComply can connect to your AWS account(s) to perform compliance scanning and monitoring. You can integrate one standalone AWS account or an entire AWS Organization (covering all member accounts).

## Prerequisites Checklist

- **AWS Account Permissions:**
  - **Single Account:** Administrator privileges in that account.
  - **Multi-Account (Organization):** Administrator or delegated admin in the AWS Organization’s Management (Root) account.
- **AWS CLI:**
  - Installed and configured on your local machine.
- **OpenComply Account:**
  - Access to the OpenComply Dashboard.

## Integration Methods

### 1. AWS Single Account (Console)

*Ideal for users with a single AWS account preferring GUI-based setup.*

#### Prerequisites

- Administrator privileges in the AWS account.
- Browser access to the AWS Console.

#### Steps

1. **Log in to the AWS Console**
   - Use an account or role with permissions to create IAM users and attach policies.

2. **Create an IAM User**
   - Navigate to **IAM** → **Users** → **Add users**.
   - **User name:** e.g., `opencomplyIAMUser`.
   - **Access type:** Check `Programmatic access` to generate an Access Key and Secret Key.
   - **Permissions:** Attach the AWS Managed Policies:
     - `ReadOnlyAccess`
     - `SecurityAudit` *(Add any extra policies only if required for your compliance needs.)*
   - **Tags:** Optional.
   - **Review and create** the user.

3. **Retrieve Access Keys**
   - On the final screen (Create User wizard), note/copy/download the `AccessKeyId` and `SecretAccessKey`.
   - **Important:** You cannot retrieve the `SecretAccessKey` again, so store it securely (e.g., a password manager, vault, or encrypted file).

4. **Add AWS Account to OpenComply**
   - Log in to the OpenComply Dashboard.
   - Navigate to `Integrations` → `AWS` → `Add AWS Account` (or similar wording).
   - Provide the `AccessKeyId` and `SecretAccessKey` from the IAM user.
   - *(Optionally)* Give it a nickname (e.g., “SingleAccountProduction”).
   - Complete the wizard to let OpenComply start scanning.

#### Validation

- Confirm OpenComply is receiving data from your AWS account (e.g., resources, compliance checks).
- Ensure no permission-related errors are reported.

---

### 2. AWS Single Account (CLI)

*Best suited for single-account setups where automation via the AWS CLI is preferred.*

#### Prerequisites

- The AWS CLI is installed and configured on your machine.
- You have Administrator permissions in the target AWS account.

#### Steps

1. **Configure AWS CLI**
   ```bash
   aws configure
   ```
   - Ensure your local credentials have administrator privileges.

2. **Test Configuration**
   ```bash
   aws sts get-caller-identity
   ```
   - It should display your account ID and ARN.

3. **Create an IAM User (CLI)**
   ```bash
   aws iam create-user --user-name opencomplyIAMUser
   ```

4. **Attach Managed Policies**
   ```bash
   aws iam attach-user-policy \
     --user-name opencomplyIAMUser \
     --policy-arn arn:aws:iam::aws:policy/ReadOnlyAccess

   aws iam attach-user-policy \
     --user-name opencomplyIAMUser \
     --policy-arn arn:aws:iam::aws:policy/SecurityAudit
   ```

5. **Create Access Keys**
   ```bash
   aws iam create-access-key --user-name opencomplyIAMUser
   ```
   - Save the returned `AccessKeyId` and `SecretAccessKey` securely.

6. **Add AWS Account to OpenComply**
   - Log in to the OpenComply Dashboard.
   - Go to `Integrations` → `AWS` → `Add AWS Account`.
   - Provide the `AccessKeyId` and `SecretAccessKey` generated in the previous step.
   - *(Optionally)* name this integration (e.g., “SingleAccount-CLI”).
   - Complete the wizard.

#### Validation

- OpenComply should start ingesting resources from this AWS account.
- Check the UI for any permission errors.

---

### 3. AWS Multi-Account / Organization (CLI)

*Designed for organizations managing multiple AWS accounts using AWS Organizations.*

#### Prerequisites

- AWS CLI is configured with Management (Root) account admin credentials (or delegated admin).
- Ability to create StackSets, IAM roles, and IAM users in the Management account.

#### CloudFormation Template

Save the YAML below as `AWSOrganizationDeployment.yml`. **Do not apply it to member accounts. It’s for the Management account:**

```yaml
AWSTemplateFormatVersion: '2010-09-09'
Description: Deploys OpenComply Platform to AWS Organization, targeting only Organizational Units (OUs)
Parameters:
  IAMUsernameInOrganizationAccount:
    Type: String
    Default: OpenComplyIAMUser
    Description: IAM User to create
  RoleNameInAccount:
    Type: String
    Default: OpenComplyReadOnly
    Description: The name of the role that will be assumed in each member account.
  OrganizationUnitList:
    Type: CommaDelimitedList
    Description: >
      List of Organizational Unit (OU) IDs to deploy the stackset to.
      Enter each OU ID without spaces.
  RootID:
    Type: String
    Description: The root ID to deploy to if no OU is specified.
    # Replace r-root1234 with your *actual* root ID, e.g. r-abcd
    Default: r-root1234
Conditions:
  HasOUs: !Not [ !Equals [ !Join ["", !Ref OrganizationUnitList ], "" ] ]
Resources:
  # IAM Role in the Management Account
  OrganizationRole:
    Type: 'AWS::IAM::Role'
    Properties:
      RoleName: !Ref RoleNameInAccount
      Description: Allows the OpenComply platform to gather inventory of the organization and member accounts
      AssumeRolePolicyDocument:
        Version: 2012-10-17
        Statement:
          - Effect: Allow
            Principal:
              AWS: !Sub '${AWS::AccountId}'
            Action:
              - 'sts:AssumeRole'
              - 'sts:TagSession'
      Policies:
        - PolicyName: OpenComplyRoleAssumption
          PolicyDocument:
            Version: "2012-10-17"
            Statement:
              - Effect: Allow
                Action:
                  - 'organizations:List*'
                  - 'sts:AssumeRole'
                Resource: '*'
      ManagedPolicyArns:
        - arn:aws:iam::aws:policy/SecurityAudit
        - arn:aws:iam::aws:policy/ReadOnlyAccess
        - arn:aws:iam::aws:policy/AWSOrganizationsReadOnlyAccess
        - arn:aws:iam::aws:policy/AWSBillingReadOnlyAccess
        - arn:aws:iam::aws:policy/AWSSSODirectoryReadOnly
        - arn:aws:iam::aws:policy/AWSSSOReadOnly
        - arn:aws:iam::aws:policy/IAMAccessAdvisorReadOnly
        - arn:aws:iam::aws:policy/IAMAccessAnalyzerReadOnlyAccess

  # IAM User in the Management Account
  IAMUserInOrganizationAccount:
    Type: 'AWS::IAM::User'
    Properties:
      UserName: !Ref IAMUsernameInOrganizationAccount
      ManagedPolicyArns:
        - arn:aws:iam::aws:policy/ReadOnlyAccess
        - arn:aws:iam::aws:policy/SecurityAudit
        - arn:aws:iam::aws:policy/AWSOrganizationsReadOnlyAccess
        - arn:aws:iam::aws:policy/AWSBillingReadOnlyAccess
        - arn:aws:iam::aws:policy/IAMAccessAnalyzerReadOnlyAccess
        - arn:aws:iam::aws:policy/IAMAccessAdvisorReadOnlyAccess
      Policies:
        - PolicyName: OpenComplySSOPermissions
          PolicyDocument:
            Version: "2012-10-17"
            Statement:
              - Effect: Allow
                Action:
                  - 'sso:Describe*'
                  - 'sso:Get*'
                  - 'sso:List*'
                  - 'sso:Search*'
                  - 'sso-directory:DescribeDirectory'
                Resource: '*'
        - PolicyName: OpenComplyAssumeRolePolicy
          PolicyDocument:
            Version: "2012-10-17"
            Statement:
              - Effect: Allow
                Action:
                  - 'sts:AssumeRole'
                Resource: !Sub 'arn:aws:iam::*:role/${RoleNameInAccount}'

  # StackSet to Deploy Roles to Member Accounts (OUs)
  MemberAccountRoleStackSet:
    Type: 'AWS::CloudFormation::StackSet'
    Properties:
      StackSetName: OpenComplyMemberAccountRollout
      Description: Stack Set that will roll out to member accounts within specified Organizational Units
      Capabilities:
        - CAPABILITY_NAMED_IAM
      PermissionModel: SERVICE_MANAGED
      AutoDeployment:
        Enabled: true
        RetainStacksOnAccountRemoval: false
      ManagedExecution:
        Active: true
      StackInstancesGroup:
        - DeploymentTargets:
            OrganizationalUnitIds: !If
              - HasOUs
              - !Ref OrganizationUnitList
              - !Ref RootID
          Regions:
            - !Ref AWS::Region
          Parameters:
            - ParameterKey: OrganizationIAMUserArn
              ParameterValue: !Sub "arn:aws:iam::${AWS::AccountId}:user/${IAMUsernameInOrganizationAccount}"
            - ParameterKey: MemberAccountRoleName
              ParameterValue: !Ref RoleNameInAccount
      TemplateBody: |
        {
          "AWSTemplateFormatVersion": "2010-09-09",
          "Description": "Create a reader role in member accounts.",
          "Parameters": {
            "OrganizationIAMUserArn": {
              "Type": "String",
              "Description": "The IAM User ARN that is allowed to assume the role."
            },
            "MemberAccountRoleName": {
              "Type": "String",
              "Description": "The name of the role that will be deployed in each member account."
            }
          },
          "Resources": {
            "OpenComply": {
              "Type": "AWS::IAM::ManagedPolicy",
              "Properties": {
                "ManagedPolicyName": "OpenComplyPolicy",
                "Description": "A Limited policy to allow cloudquery to do its job",
                "PolicyDocument": {
                  "Version": "2012-10-17",
                  "Statement": [
                    {
                      "Effect": "Deny",
                      "Resource": "*",
                      "Action": [
                        "dynamodb:GetItem",
                        "dynamodb:BatchGetItem",
                        "dynamodb:Query",
                        "dynamodb:Scan",
                        "ec2:GetConsoleOutput",
                        "ec2:GetConsoleScreenshot",
                        "ecr:GetAuthorizationToken",
                        "ecr:GetDownloadUrlForLayer",
                        "kinesis:Get*",
                        "lambda:GetFunction",
                        "logs:GetLogEvents",
                        "s3:GetObject",
                        "sdb:Select*",
                        "sqs:ReceiveMessage"
                      ]
                    }
                  ]
                }
              }
            },
            "MemberAccountReadOnlyRole": {
              "Type": "AWS::IAM::Role",
              "Properties": {
                "RoleName": { "Ref": "MemberAccountRoleName" },
                "Description": "Read Only Access to fetch inventory from member accounts",
                "ManagedPolicyArns": [
                  { "Ref": "OpenComply" },
                  "arn:aws:iam::aws:policy/ReadOnlyAccess",
                  "arn:aws:iam::aws:policy/SecurityAudit",
                  "arn:aws:iam::aws:policy/AWSOrganizationsReadOnlyAccess",
                  "arn:aws:iam::aws:policy/AWSBillingReadOnlyAccess",
                  "arn:aws:iam::aws:policy/IAMAccessAnalyzerReadOnlyAccess",
                  "arn:aws:iam::aws:policy/IAMAccessAdvisorReadOnly"
                ],
                "MaxSessionDuration": 28800,
                "AssumeRolePolicyDocument": {
                  "Version": "2012-10-17",
                  "Statement": [
                    {
                      "Effect": "Allow",
                      "Principal": {
                        "AWS": { "Fn::Sub": "${OrganizationIAMUserArn}" }
                      },
                      "Action": [
                        "sts:AssumeRole",
                        "sts:TagSession"
                      ]
                    }
                  ]
                }
              }
            }
          }
        }
Outputs:
  IAMUserNameInMasterAccount:
    Description: IAM Username in the Master Account.
    Value: !Ref IAMUsernameInOrganizationAccount
  IAMRoleName:
    Description: IAM Rolename that is created.
    Value: !Ref RoleNameInAccount
```

#### Step-by-Step Deployment

1. **Log in to AWS (CLI) using the Management (Root) account**
   ```bash
   aws sts get-caller-identity
   ```
   - Confirm it shows the Management account ID.

2. **Retrieve Your AWS Organization Root ID**
   ```bash
   ROOT_ID=$(aws organizations list-roots --output text --query 'Roots[0].Id')
   echo "AWS Organization Root ID is: $ROOT_ID"
   ```
   - This typically looks like `r-xxxx`.

3. **Create the CloudFormation Stack**
   ```bash
   aws cloudformation create-stack \
     --stack-name opencomply-Deploy \
     --template-body file://./AWSOrganizationDeployment.yml \
     --capabilities CAPABILITY_NAMED_IAM \
     --parameters ParameterKey=OrganizationUnitList,ParameterValue=$ROOT_ID
   ```
   - If deploying to specific OUs, provide them comma-separated:
     ```bash
     aws cloudformation create-stack \
       --stack-name opencomply-Deploy \
       --template-body file://./AWSOrganizationDeployment.yml \
       --capabilities CAPABILITY_NAMED_IAM \
       --parameters ParameterKey=OrganizationUnitList,ParameterValue=ou-abc1-12345678,ou-def2-87654321
     ```

4. **Monitor Stack Deployment**
   ```bash
   aws cloudformation describe-stacks --stack-name opencomply-Deploy --query "Stacks[0].StackStatus" --output text
   ```
   - Wait until it reaches `CREATE_COMPLETE`.
   - If there are errors, check stack events:
     ```bash
     aws cloudformation describe-stack-events --stack-name opencomply-Deploy
     ```
     - Or use the CloudFormation Console.

5. **Generate IAM Access Keys**
   ```bash
   IAM_USER=$(aws cloudformation describe-stacks \
     --stack-name opencomply-Deploy \
     --query "Stacks[0].Outputs[?OutputKey=='IAMUserNameInMasterAccount'].OutputValue" \
     --output text)
   echo "IAM User created by stack: $IAM_USER"

   aws iam create-access-key --user-name "$IAM_USER"
   ```
   - Save the `AccessKeyId` and `SecretAccessKey` securely.

6. **Add AWS Account (Organization) to OpenComply**
   - In the OpenComply dashboard, go to `Integrations` → `AWS` → `Add AWS Account`.
   - Provide the access keys from the previous step.
   - OpenComply will assume the roles in each member account for read-only scanning.

#### Validation

- Check OpenComply to confirm data is pulled from each member account in your organization.
- Verify there are no permission errors in the UI.

---


## Troubleshooting

- **Access Denied Errors:**
  - Ensure the IAM user has the necessary permissions (`ReadOnlyAccess` and `SecurityAudit`).
  - Verify that the Access Key and Secret Access Key are correctly entered in the OpenComply Dashboard.

- **CloudFormation Stack Deployment Issues:**
  - Check the CloudFormation Console for detailed error messages.
  - Ensure that the AWS CLI is configured with the correct Management (Root) account credentials.

---

## Additional Notes

- **CloudFormation Template:** Ensure that you replace placeholder values such as `r-root1234` with your actual AWS Organization Root ID.
- **Permissions:** The CloudFormation template includes various managed policies to ensure OpenComply has the necessary read-only access across your AWS Organization.
- **Automation:** Using the CLI and CloudFormation enables automated and scalable deployments, especially beneficial for large organizations with numerous AWS accounts.

By following this SOP, you can efficiently integrate OpenComply with your AWS environment, ensuring comprehensive compliance monitoring and scanning across single or multiple accounts.

---
