# Continuous Delivery/Deployment

Created 2 workflows for the project. 
One for the `main` branch and one for the `releases/**` branches. 
Both workflows under the hood use the same action that builds, tests, and deploys the application to the provided environment.


It's type of continuous delivery, not deployment as it's expected that after deployment on stage environment, someone will ensure everything is working and then deploy to production by creating release branch.