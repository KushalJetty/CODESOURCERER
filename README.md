# **CODESOURCERER**  
### *An AI-Driven Automated Test Generation Tool*  

<picture>
  <source srcset="https://github.com/user-attachments/assets/32f0cd5f-9774-4af8-84b6-746ff03f74de" media="(prefers-color-scheme: dark)" style="filter: invert(1);" />
  <img src="https://github.com/user-attachments/assets/b350c464-6b9e-4d93-babc-2914f6a34e3b" alt="description" width="200" align="center" style="filter: invert(0);" />
</picture>


#### To know more visit:  [Website](https://codesourcerer.webflow.io/)

#### Proposal: [Slides](https://docs.google.com/presentation/d/1bkRmrLHOkwKDrVaksg7pQV5RSit0jxG1Tvpl_AGFNko/edit?usp=sharing)
---

## **Overview**  
CODESOURCERER is an innovative solution designed to streamline the software development lifecycle by automating the generation of test cases for code changes in GitHub repositories. Built with a robust microservices architecture using **Golang**, **Docker**, **Kubernetes**, and **gRPC**, it integrates seamlessly with CI/CD pipelines to enhance efficiency, reliability, and collaboration.

**Key Features**:  
- **Automatic Test Generation**: Uses AI (powered by a fine-tuned LLM) to create unit and regression tests.  
- **Scalable Architecture**: Microservices ensure flexibility and reliability.  
- **Developer Collaboration**: Draft Pull Requests (PRs) for reviews and notifications.  
- **Retry Mechanisms**: Automated retries for failed tests with detailed notifications.  

---

## **Why These Technologies?**  

### Golang  
- High concurrency and performance.  
- Ideal for handling microservices and large-scale requests.  

![Golang Logo](https://www.google.com/imgres?q=golang%20logo&imgurl=https%3A%2F%2Fw7.pngwing.com%2Fpngs%2F566%2F160%2Fpng-transparent-golang-hd-logo-thumbnail.png&imgrefurl=https%3A%2F%2Fwww.pngwing.com%2Fen%2Fsearch%3Fq%3Dgolang&docid=k5kORhOaFqiZ7M&tbnid=ues55JIB2kXJSM&vet=12ahUKEwi-iL7PkdOKAxWZzzgGHY3rMpIQM3oECGIQAA..i&w=360&h=180&hcb=2&ved=2ahUKEwi-iL7PkdOKAxWZzzgGHY3rMpIQM3oECGIQAA)  

### gRPC  
- Efficient communication protocol for microservices.  
- Supports bidirectional streaming and low-latency data transfer.  

![gRPC Logo](link_to_grpc_logo)  

### Docker & Kubernetes  
- Enables containerized deployment and scalability.  
- Kubernetes automates scaling, load balancing, and fault tolerance.  

![Docker Logo](link_to_docker_logo)  
![Kubernetes Logo](link_to_kubernetes_logo)  

---

## **System Architecture**  
![System Architecture Diagram](link_to_architecture_diagram)  
*(Add a visual of your architectural design)*  

---

## **Installation Instructions**  

### Prerequisites  
- **Operating System**: Windows 10+ or Linux.  
- **Software**: Golang 1.20+, Docker, Kubernetes, Redis.  

### Steps  
1. **Clone the Repository**  
   ```bash
   git clone https://github.com/your-repo/codesourcerer.git
   cd codesourcerer
   ```

2. **Build the Services**  
   ```bash
   docker-compose build
   ```

3. **Start the Services**  
   ```bash
   docker-compose up
   ```

4. **Set Up GitHub Webhook**  
   - Configure your GitHub repository with the webhook URL provided by CODESOURCERER.  
   - Use the `codesourcerer-config.yaml` to define branch triggers.

5. **Verify Setup**  
   Test the workflow by pushing code changes to the specified branch.  

---

## **Usage**  
1. Push code to the specified branch in your GitHub repository.  
2. Let CODESOURCERER handle test case generation automatically.  
3. Review the Draft Pull Request with generated tests.  
4. Merge after validation.

---

## **Contributors**
