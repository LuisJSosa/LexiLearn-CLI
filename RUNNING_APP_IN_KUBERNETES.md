# Running the LexiLearn Application on Kubernetes

This guide provides instructions for running the LexiLearn application on Kubernetes, both for local execution and for first-time setup using Google Cloud Platform (GCP).

## Table of Contents

1. [Running Locally](#running-locally)
   - [Step 1: Access the Application Pod](#step-1-access-the-application-pod)
   - [Step 2: Run the LexiLearn Application](#step-2-run-the-lexilearn-application)
2. [First-Time Setup on GCP](#first-time-setup-on-gcp)
   - [Step 1: Create a Google Cloud Account](#step-1-create-a-google-cloud-account)
   - [Step 2: Install Required Tools](#step-2-install-required-tools)
   - [Step 3: Authenticate with the Kubernetes Cluster](#step-3-authenticate-with-the-kubernetes-cluster)
   - [Step 4: Verify and Access Running Pods](#step-4-verify-and-access-running-pods)
   - [Step 5: Launch the LexiLearn Application](#step-5-launch-the-lexilearn-application)
3. [Troubleshooting and Support](#troubleshooting-and-support)

## Running Locally

### Step 1: Access the Application Pod

To interact with the LexiLearn application locally, first access the application's pod. Execute the following command in your terminal:

\```
kubectl exec -it lexilearn-deployment-848f949fb8-r8v5l -- /bin/sh
\```

This command opens an interactive shell session inside the running container, allowing you to execute commands within it.

### Step 2: Run the LexiLearn Application

Inside the pod, start the LexiLearn application:

\```
./lexilearn
\```

You can now interact with the application directly from the terminal, including adding new vocabulary, reviewing words, and managing decks.

## First-Time Setup on GCP

### Step 1: Create a Google Cloud Account

If you don't have a Google Cloud account, [create one here](https://cloud.google.com/). This is necessary for accessing Kubernetes services on GCP.

### Step 2: Install Required Tools

Ensure you have the following tools installed on your machine:

- [Google Cloud SDK](https://cloud.google.com/sdk/docs/install) 
- [Kubectl](https://kubernetes.io/docs/tasks/tools/)

### Step 3: Authenticate with the Kubernetes Cluster

Once your account is set up, authenticate your local machine with the Kubernetes cluster. Run the following command in your terminal:

\```
gcloud container clusters get-credentials lexilearn-cluster --zone us-central1-a --project lexilearn-1
\```

This command fetches cluster credentials for `lexilearn-cluster`, located in the `us-central1-a` zone and part of the `lexilearn-1` project. It configures `kubectl` to interact with the cluster.

### Step 4: Verify and Access Running Pods

After authenticating, list the running pods to verify that your application is active:

\```
kubectl get pods
\```

Look for a pod named `lexilearn-deployment-848f949fb8-r8v5l` and confirm its status is `Running`.

### Step 5: Launch the LexiLearn Application

Access the application pod and run the application with the following commands:

\```
kubectl exec -it lexilearn-deployment-848f949fb8-r8v5l -- /bin/sh

./lexilearn
\```

## Troubleshooting and Support

- **Pod Not Running:** If the LexiLearn pod is not running, check the Kubernetes cluster for any issues and restart the deployment if necessary.
- **Authentication Issues:** Ensure you have the correct permissions and that your Google Cloud SDK is up-to-date.
- **Application Errors:** If the application fails to start, check the logs within the pod for error messages by running:

  \```
  kubectl logs lexilearn-deployment-848f949fb8-r8v5l
  \```

This document is designed to help users set up and run the LexiLearn application efficiently, whether locally or on GCP for the first time. If you have any questions or need further assistance, feel free to reach out to support.
