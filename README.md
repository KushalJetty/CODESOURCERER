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

![Golang Logo](data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAT4AAACfCAMAAABX0UX9AAAAq1BMVEUArNf////v7+/u7u739/fy8vL8/PwAqtYAqNUArtio3e6j3u4YsNn9+ff18fBywt6T2OphwN9vy+Wl2OzC6fObz+N9yuPj7/Hs+PsApNQAsdmTzeI/ttug1ufh8/n//PrT7fKy4/BKt9rg8PIzvN7f7O/H5/Ps9/v2+/3T5+yCz+bR7PSU2ep1z+Zhx+NLvd6w1+bA3+qq1+c1vt6IyuPF4Oq54e7J7POD1OiX1x9+AAAQh0lEQVR4nO1di2KivBIGuSRUSdVaReQiikqhxV62tu//ZIcgakQkAYltz+/sdt2FHb+Zj5DrZCKIpCitg0gyeUeVyFs3pUyE323fL1e60Xej70bfH1W60Xej70bfH1W60Xej70bfH1W60dcofRIhOS1SbkqZCAopKiny+Vs3pUwEsihKKsmzTBRTSTl6ONJNaSuM9LXqQP0HlG703ei70fdHlX4ZfVK9Cv1GH6auhVzk6uLc+BqkMjXmCOl6eou/eVINpV9C35a60aAbx7OZZVmdVJK/zOL+c6SIKC2YXM1L8NEW5s/Rp4to+hybgiAAAPAPKclVvz/X3abNQwhf0dvtdreP5aXdHrRQQuMfoy+Ytn0IYUrUGYFjrb9ZiZciZUqoFQSjwff3hzAej+GRjE2/661GwV+hz3BefQGWMJcJgNB8naoXIKVKiowWT9Hbi4m5KkIFAoSC342cgAVJOBoB52ukg+Sr2YaUomVfg1Tq9gwKdqTWRNqKvvp67SePi/a8kmcVLh06kiD/mIjBoOtTHckRaE6WoloTT5cHD0OTETB5Vj37SxbLv/PHJqxc1L63qlCXeQWs7khvVTdPRl92AlhWwRZg3Q+QXDphdQxFVhpc5yP1dk+r4goh2sxD6gl9peZJsj69t1gLHiHQuo908SzST802T2uTlwjQ1oZbwTwJuZ/rTqWCd8ACnW5wFuln6DNCrTZ3GYEeYjdPfPJLe0UULGgufxN96rLGW3TCX6jQkVIJFh/jC+HGE+PX0OdMLvUmFXg/YjLP6dYveAcwP/od9LlLn7mfR3GpF9HNM+7WTTwrAZhL9RfQZ7ybjbiDBVoDmnnT4YWVLCG28eP0OX2hMXeSImF9lZvnzZqDA0Lf+WH6nLiBeoh0ac9foXn9TpNoQPCdH6Vv1dyLu3NpZpw1z5g1+6xwBbiiRRnUoY9JSRKfxs16gwX25DPmObOGmihCgLbK03d+pV1mXIRnVNpwYC+RsNA8cWM1zx7mz3CPkK42YbXh4A12qHMnt07Mk+8arycyuPXiCOlK06XuptFqnHTIekJ581q82EskNkikK9G37PByJ6n+VnnzuvzYA0KX9Pc69EU1ZvbYPWqrxy9vW+MHllR/m2vT51gc/UkKhHJE35fJEy2p/pwD1jXoMz44lr1EYCwS5g24PisM1z9Uf1egT+3z6EOQMia6mzz6e3m4w/zfFeiz+XT4CAHxHj8YcmdPEMz961tOn9TKok9O+n3Z5WKlI/o+q/mzHWlVG28BbTf3EnT51hNbgSELfTLS8a/0D3RE3+Fy8uGSd9LLJH2RX2VtC0LT9BMxzfKog7zeMIP3KnZZwD68oNIDA7var4Q+ffncttupZB972V/O39heRAf6gj6zWQAKcfctiqb493L5HrPP6e+mXkZVOucJdcK6b9vL5du7Ha7NCkvO8IVO36ZTUzSCvgfWGTcA/XdnJKH9EpBqOMuY1SOtm+qs2SsKOF63HWdlBHgdTVWDlfPZZ35eydg3o+/s8NXxQT0ZE2PeL8b+MhjHUSC3WrmBsip+jFniX7La6IWRPQDHvjcKciM9SZFtjZXArLE6F2XQMuKaTRh83n+d+8RWHIAWtw6U5QIGQq20ACd1ltaxuiNJcqdsFgOtsx14FYUmDDpMrwuwpq3SKIPXuuytD6vK6iPLlwBhTY6DTqbujOFZjwDoWLPQ03FXSmGr+ACYtYMzSKm8MH2N1tWx9jn6Po9fSNwyAXhop2D2DwCJixrUNMEiZnQilnYQdCZHkxinTgUPBR4l5ghWb9idu1lBkkOWoS60QuM8UirejOF7QK+llNDXXvd6vWFvuJNJ+jPJ5IGU9Mrj4+TRfrBtm1w9ZBluwPVGOer6FDm1yX8RBGZvYn8RTKABQzULwP2yHAnLImb4po6Hi985+hZPU8cxnGAn8sgwlBHK5EhJTqOCt91ARMbTOAxlD8ZTxLAs8EqMXJJi7tuvA+dISVkx9JCAFho0pNTwmG44mOglpa/VQke9GFVSlEIo3N4Qd0gmGB7ieGIgllUVd5KWP1yB+OFb5LSQfGweemVY0e20z0f7HPm0oi+ug96ilL6LV9oG9Fd3/K4qEhOS4SeV7NiM7YWhtPCTzCkt6BUF0CKV1acWtf0F/kDmSh99tAbfETOS8R17yb/OmOfSOwoAzlVmn5BH/T4NG8+PvjbVHzAJlGpIxNaVYyWDWlEAYSpWiKkUqaUZPnKlr0ctfMNgu0TcwIL8kjYpBoS2ylZNZA9KoXX/wHCUpw83nnIy7Exb0aPuhJg1rnL6cUaJdIo6XANrJ5scvJy+gN5KhUG1cq7e0R6IGaGUPnVwtxMb/7K3H3fHsrtccGd3eUTYF9L8ESJdaqacM3SR4P0IVUSitb47+tSNJWhNiOAdunAG7d2Fobsfbl4cwx/SKiorSqPxKyC1lHfKl4I3JAkK+mxoBRvAYJ82AdmU8RrQArGxBA06tZvRFSsjubQhJ7RVVUBGhUmyckZMMVvrkJRgSHFo7BHPuiTXAnJ1XUdi8iv5PLqju/hyepfWywDrOR3pZP2GNuGR0CcKzJNkVIHtvX1oQyt8a7L/TzolLrwuKc/7P87KM6WHCTrtIiTq8temfA4CL1gKd+Oas6InApW9fTKt3oDkFBXhlCL1c1scWaQc6xAFWJG+aE2nr98j5J6UXo/t1u7i8GDfiFaX+04hfZIb8ljV7BaWcyp9lBoIxAl9aDQyst/GSNYJUQ7Xk9/kHX1uFCgd7FOm5c9NgJPiDbNy1CRrOzdNpr2lp/Rl8xRnv9c30qUiRZGS0a+cfOb2viotWcI3k8/cusD2epEStgJ90lrCN/FECVutM0zSVBb4XfCgWOgT7fLSZzqsUQYV81tQqr6jKBsCSR7wiAIcj06R2OgrDxOsQF+1II2AUvXBcFSEJCFad6eOAK22T+WL7tzoW1BCquAjKkLSad2dWkK+u3+DvhVl5lfz9CIk5R+P8B4YnSL9bvqeKFWfNZALkNAnpb2uJ/BcB70R+moEyaezVwcR80qtk4WxHG7PyTfXWEsq7yfUFPAR5JGKfSoggtJ04I4LYx6FlUeuTT6Ssr20MQ7/uU2hb6gUIOhTv1nitgLDmpkjEosmlG7zijkVxJ1GGTjFjrJX+qY0vP+KkNQ2l9A8aMvnfKJNjaHyuhj2A9bJ+oC27gPi+UGJ0vlN6TtBMrgUPgF653yiTczSpo1gf8RKH+V1xItmhBKl34KjyU6RlnyiauvT13oqn/KFtsRK37lpGfw16efRmiOl+1ZMH6etLPXpkz3KhJXNutL2sot8tAiZkbJ+u5A+HvstUydr04do06VviJG++SCTr2k034mczbLgcGZXlysUpCL6qoRAV5H69I0o81W7lTY6faKIo4C2K5WnUDgqtMp7WNR0MIY2Vpek5a1HH6KtdfhTdvpSqN2mA9p2aCp97gkSlwEHFhAHNUsfZQtt0tdQ+EQZ0BY6hkFeyeC1XzUZG9Skb0WJToaPLU5BGpSOSzJoyyvNudEnwFE9nxYUizRP/yH6NC+vpHOk76uOTxKt7w/MggC1k2qMCSqvRBm0CfDhRKnbXL6avJ8fh8kUdp+UBaUjBYYjiVMqiC512dqRc0qi1y+V+k0LMMUaPlE7UmCic0oFQZuwSuRomTdVSsMJ9DTtNSlZlIHIvrfrROBTdZ/ogQvbDFBcpkupmyjJHcU7pZK9m5JLjQEq83Rd1SdJpm6pAD2VG30ObYcF0BYVkCTdvqhTrQ1YkTKRHeqCqbbdk8pnqYg69wT7KjOSpA8um04APZcNaX+DGpwIrCk/+gKGMPcpMxKixSxQsTpeNZ8m9Lj6rELgsxn/jk7fmBFJko37S8fDwIoYkHaXEcN2PuBxpA8xbMAHvsqEpEgNLF/Ce5nZJ6lLnzoDHZEjfQrLXkDQDxiQWkEz629Z7h+qT/KCgT0B7hojPvTNmTbVdQ060hM9MwZg2EkJhDaDT5KCpkxZU6ydEh/6ZJaNvEDrr2hIDj1yGJjfLPyZXapPEgo2TJHKh0QufJoOanBuxl+8KEVy7xi8gTZ1B0sKJoQUnxT9aci2o/oQWcyHPmV0z7az22qXIDkseW+BGYlMS5xAuDdKD9tRuxbbRnyNyCLEYcyLR7ATFo+wKb5xBimgJDHIBD4i9MyYxMDyjqCOfZr6jEk/YHgI++B14MSUNd00gP7mJH+6GkgvLCk0tjPXosi4SgfgehqQSKkk5VBW7kzWlT7gR2jvLrctgUNGa7BTZjdyjEANAjX5MVZPS3wMCpvyNv8N8+YKMPY/98fo4C3eCkLBwnlnX+cDwjs6+MuNvipxtgCONT+0bfs1+enHgJW7ZOT+kIK1Kky2QjN8X0YRzjMgGdMoerNjlrOS9tJ3CIe50efSN6SSAohsUhW04m3Fq1aZ0MJPx/R74b9//4Zrv0ruJSF7da9An/hUwaiaAvaZ4KKKUzJgm/wLHwpXUe+TeHW5bsbnscngWMb7AP2rpO1LBHZdkgie9AW80x6O7w5gS37JcgmBH1fMWU9bcLvUF6L/pdLieRoRvD/pevSpvMJ+tr6sifUSVXaqNVW1EC3jqicmDJo8nSPvy9HOJFVCr/xW2jPEzlOeCL70BSG/7N1WdGyerP/jhbVD3JwQwZc+cdTjVCOBzqeeN89p5liic7Jlj3JeR1Nj3q0Sms64uASEjX5iHnI4tr5As4uIEJhyLdRK0JAIToDKw5nxIt0vmzNCdLgFyiQdvqDIXT4bUgklj4NLmL1C8xxeXc3xl1hIBP+s4d/NHxlknDXv8+LzBwsRwdfe5aufDv3daPcF7PIgFJqnvjfffQHaenDw+PqHa383+P4Cre+Umee+Nc0fEPoLYpLlB84m7zZ32olmU1Y33bdmj1YBna6DCpGuRp84aOgUCLg/KrLEvKiptEg7xLKj3a9Cn7hogj+cWpzFPKepU0QTRFxV/Dx9IkPQFU2g8Mq4MXzV0BEhECyDcqRr0Yc3ZV4WpAfjJ1bzFPHTv/wEZbCbEfsV9IlTv/4JVjgxu1uUAKHYPEl3+pc19wBoPYeOdEX63FabbRm/QKzQ0E8W5EvMk2Td69UnEMBOz2NButohi1kQiVl9DAIEa/jlVjZPHzEGXRSRd99mQ2JNBdGUDCbVznUHUOvZX7JYB0uM7F6NpwX80At0Ngi+E1YFSsHynf1QFgDNcOnURMLieHGlEgjgeH0XBcxI1z7aHV91NmvAcgYM1Pw7J7gACYuxicfMh8hArR85FZCufbR7JoH0YpaGAOFF7HXCXetSpEQpWHj+mH4SVspdEFRC+iH6EkGt57WgCSe5JfDhOZoQ9+e62wxSoqS4o/ZHRzhkryCe0+6a/zKvjvRj9OHd7q7y9RzHM8uy9okmZrO4/+0pKO18NIaUfJUu6vOXeI2hdpzhmGgNo87W/UEul9QfoC/5z64rIjQfTbNEE9M5VtdlpbXL7d8QUiv9QtFFSjRoh2GWbDUMw643GIzwWX31kH6UPuxU8j/kayBlaIcAMPwH3sGJM+3XRvph+jKvroV0qoR3I/xx+v6w0o2+G32/hz6JSYuWCuI/o8QnFcR/R+maE1b/h0pXnC79f1S60Xej70bfH1W60Xej70bfH1W60Xej70bfH1W60Xej7+eU/gdHR/5Et6LUsAAAAABJRU5ErkJggg==)  

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
