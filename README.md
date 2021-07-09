# template-go  

This repo is a demo/template repo that demonstrates how to setup a repository that will run github actions to lint the code, run unit tests, and run code coverage reports.  The following steps need to be completed in order to create a new repo from this template.

    1) On the main page of this repo, click the 'Use this template' button
    2) Fill out the form, ensuring that you mark the new repo as private, and do NOT 'Include all branches'
    3) In the new repository, navigate to settings -> manage access and invite the maxnext team
    4) Next, navigate to settings -> Branches and ensure that main is the default branch, then create a branch protection rule and enable at least the following:
        a) Type 'main' as the branch name pattern
        b) Require pull request reviews before merging and choose 2 as the number of approvals needed
        c) Require status checks to pass before merging, Require branches to be up to date before merging, and choose the status checks that coorespond to your test suite.
        d) Include administrators
        e) Save changes


The github actions workflow file can be extended to do many other types of golang specific actions.  Search [here](https://github.com/marketplace?type=actions) for other github actions.


General workflow to achieve ci/cd is as follows. 

    1) Code is commited to a non-main branch
    2) A PR is created to merge that branch into main
    3) Once the PR is created and any subsequent pushes will trigger github actions to run which will:
        a) Lint the go code
        b) Run unit tests
        C) The results of the linting and unit tests will either block or allow the ability to merge the PR.
    4) Once the PR is approved and merged, it will trigger the ci/cd pipeline located in aws codepipeline.

Golang Files included in this template are for demonstration purposes. The other workflow files are described below:

**.github/workflows/test.yaml** - The configuration file for the github actions workflow.  The github action will typically run when a PR is created, when a commit is made to a branch with an open PR, and when that PR is merged to the main branch. This behavior is controlled by the "on" statement at the top of the file.

**.taskcat.yml** - The configuration file for taskcat which runs during the ci/cd pipeline. More information on taskcat can be found [here](https://github.com/aws-quickstart/taskcat).

**buildspec.yml** - Defines aws codebuild job that is run during ci/cd pipeline (this is where taskcat runs).

**cloudformation_template.yaml** - A cloudformation template that creates the CloudFormation role that CodePipeline assumes to deploy `template.yaml` located in this repo.

**codepipeline_template.yaml** - A cloudformation template that creates a CI/CD pipeline in CodePipeline. This should use the `ci-cd-pipeline-module` found [here](https://github.com/maxexllc/ci-cd-pipeline-module/).

**template.yaml** - The cloudformation template for this application.  This is where most of your changes will take place.

Please note that this pipeline can/will/should evolve over time. This template was created as a starting point.  Any of these steps and/or configurations can change per project as long as the development team agrees on the change. The general idea is to automate as much testing as possible, have a set of standards that will help us decide whether to accept a PR, and have a fully automated deployment capable of automatically rolling back in the event of an error.
