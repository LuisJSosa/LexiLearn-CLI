# Running LexiLearn in Kubernetes

This guide provides instructions for setting up and running the LexiLearn application in a Kubernetes environment. It includes steps for installing the necessary tools and accessing the application on different operating systems.

## Table of Contents

1. [Install Google Cloud SDK](#install-google-cloud-sdk)
   - [Windows](#windows)
   - [macOS](#macos)
   - [Linux](#linux)
2. [Common Setup Steps](#common-setup-steps)
3. [Access the Application](#access-the-application)
4. [Troubleshooting and Support](#troubleshooting-and-support)

## Install Google Cloud SDK

The installation process for the Google Cloud SDK (gcloud CLI) varies depending on your operating system. Follow the steps below for your platform.

### Windows

1. **Download the Installer:**
   - Go to the [Google Cloud SDK for Windows download page](https://cloud.google.com/sdk/docs/install).
   - Download the installer file for Windows.

2. **Run the Installer:**
   - Double-click the downloaded installer file.
   - Follow the prompts to install the Google Cloud SDK.
   - During installation, you can choose to install additional components like Python and the Cloud SDK component manager.

3. **Initialize the SDK:**
   - Open a new Command Prompt window and run the following command to initialize the SDK:
     ```
     gcloud init
     ```

### macOS

1. **Using Homebrew:**
   - If you have Homebrew installed, you can easily install the Google Cloud SDK using the following command:
     ```
     brew install --cask google-cloud-sdk
     ```

2. **Initialize the SDK:**
   - After installation, initialize the SDK:
     ```
     gcloud init
     ```

### Linux

#### Using a Package Manager (APT for Debian/Ubuntu):

1. **Add the Cloud SDK distribution URI as a package source:**
   ```
   echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] http://packages.cloud.google.com/apt cloud-sdk main" | sudo tee -a /etc/apt/sources.list.d/google-cloud-sdk.list
   ```

2. **Import the Google Cloud public key:**
   ```
   curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key --keyring /usr/share/keyrings/cloud.google.gpg add -
   ```

3. **Update the package list and install the Cloud SDK:**
   ```
   sudo apt-get update && sudo apt-get install google-cloud-sdk
   ```

#### Using a Package Manager (YUM for CentOS/RHEL):

1. **Create a repo file:**
   ```
   sudo tee -a /etc/yum.repos.d/google-cloud-sdk.repo << EOM
   [google-cloud-sdk]
   name=Google Cloud SDK
   baseurl=https://packages.cloud.google.com/yum/repos/cloud-sdk-el7-x86_64
   enabled=1
   gpgcheck=1
   repo_gpgcheck=1
   gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg
          https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
   EOM
   ```

2. **Install the Cloud SDK:**
   ```
   sudo yum install google-cloud-sdk
   ```

## Common Setup Steps

### Initialize the SDK

After installation, run the following command to initialize the SDK and set up your configuration:

```
gcloud init
```

This command will guide you through authorizing the SDK with your Google account and setting up the default project and configuration.

### Verify Installation

To verify the installation, run:

```
gcloud --version
```

This command should display the version information for the installed components.

### Install the GKE Auth Plugin

Run the following command as an administrator to install the `gke-gcloud-auth-plugin` binary:

```
gcloud components install gke-gcloud-auth-plugin
```

## Access the Application

### Step 1: Authenticate with the Kubernetes Cluster

Authenticate your local machine with the Kubernetes cluster using the following command:

```
gcloud container clusters get-credentials lexilearn-cluster --zone us-central1-a --project lexilearn-1
```

This command fetches cluster credentials for `lexilearn-cluster`, located in the `us-central1-a` zone and part of the `lexilearn-1` project. It configures `kubectl` to interact with the cluster.

### Step 2: Verify and Access Running Pods

After authenticating, list the running pods to verify that your application is active:

```
kubectl get pods
```

Look for a pod named `lexilearn-deployment-848f949fb8-r8v5l` and confirm its status is `Running`.

### Step 3: Access the Application Pod

Once you've identified the running pod, access its shell to interact with the application:

```
kubectl exec -it lexilearn-deployment-848f949fb8-r8v5l -- /bin/sh
```

This command opens an interactive shell session inside the running container, allowing you to execute commands within it.

### Step 4: Run the LexiLearn Application

Inside the pod, start the LexiLearn application:

```
./lexilearn
```

You can now interact with the application directly from the terminal, including adding new vocabulary, reviewing words, and managing decks.

## Troubleshooting and Support

- **Pod Not Running:** If the LexiLearn pod is not running, check the Kubernetes cluster for any issues and restart the deployment if necessary.
- **Authentication Issues:** Ensure you have the correct permissions and that your Google Cloud SDK is up-to-date.
- **Application Errors:** If the application fails to start, check the logs within the pod for error messages by running:

  ```
  kubectl logs lexilearn-deployment-848f949fb8-r8v5l
  ```

This document is designed to help users set up and run the LexiLearn application efficiently on Kubernetes. If you have any questions or need further assistance.
