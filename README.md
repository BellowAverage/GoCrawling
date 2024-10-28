## Web Scraping Wikipedia Articles on Intelligent Systems and Robotics

### Description

This project involves developing web crawlers and scrapers to extract text information from a list of Wikipedia articles related to intelligent systems and robotics. The goal is to retrieve the textual content from specified Wikipedia pages, ignoring images and HTML markup, and save the extracted text into a JSON Lines file (`articles.jsonl`).

Two implementations are provided:

1. **Go-based scraper using the Colly framework** (`main.go`)
2. **Python script using Requests and BeautifulSoup** (`wikipedia_scraper.py`)

### Performance Observations

- **Go-based Scraper Execution Time**: **359 milliseconds**
- **Python Script Execution Time**: **2.92 seconds**

**Observation**:

- The Go-based scraper demonstrates significantly faster execution time compared to the Python script. This performance gain is primarily due to Go's ability to perform concurrent processing using goroutines.
- The Python script processes the URLs sequentially, which results in longer execution time, especially noticeable when dealing with multiple URLs.
- The use of concurrency in Go allows for multiple web pages to be scraped simultaneously, efficiently utilizing system resources and reducing total execution time.

---

### List of Wikipedia URLs

The following Wikipedia pages are scraped:

- [Robotics](https://en.wikipedia.org/wiki/Robotics)
- [Robot](https://en.wikipedia.org/wiki/Robot)
- [Reinforcement learning](https://en.wikipedia.org/wiki/Reinforcement_learning)
- [Robot Operating System](https://en.wikipedia.org/wiki/Robot_Operating_System)
- [Intelligent agent](https://en.wikipedia.org/wiki/Intelligent_agent)
- [Software agent](https://en.wikipedia.org/wiki/Software_agent)
- [Robotic process automation](https://en.wikipedia.org/wiki/Robotic_process_automation)
- [Chatbot](https://en.wikipedia.org/wiki/Chatbot)
- [Applications of artificial intelligence](https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence)
- [Android (robot)](https://en.wikipedia.org/wiki/Android_(robot))

---

### Implementation Details

#### 1. Go-based Scraper (`main.go`)

- **Language**: Go
- **Framework**: [Colly](https://github.com/gocolly/colly)
- **Concurrency**: Utilizes goroutines and WaitGroups for concurrent processing
- **Execution Time**: Approximately **359 milliseconds**
- **Output**: `articles.jsonl` containing JSON Lines with the URL and extracted text

#### 2. Python Script (`wikipedia_scraper.py`)

- **Language**: Python 3.x
- **Libraries**: `requests`, `beautifulsoup4`
- **Concurrency**: Sequential processing
- **Execution Time**: Approximately **2.92 seconds**
- **Output**: `articles.jsonl` containing JSON Lines with the URL and extracted text

---

### Requirements

#### Go-based Scraper

- Go installed (version 1.13 or later)
- Colly framework:

  ```bash
  go get -u github.com/gocolly/colly
  ```

#### Python Script

- Python 3.x installed
- Required libraries:

  ```bash
  pip install requests beautifulsoup4
  ```

---

### Installation and Usage

#### Go-based Scraper

1. **Clone or Download the Repository**

   Ensure you have the `main.go` file in your working directory.

2. **Install Dependencies**

   Install the Colly framework:

   ```bash
   go get -u github.com/gocolly/colly
   ```

3. **Run the Scraper**

   Navigate to the directory containing `main.go` and execute:

   ```bash
   go run main.go
   ```

4. **Output**

   The extracted articles will be saved in an `articles.jsonl` file in the same directory.

#### Python Script

1. **Clone or Download the Script**

   Ensure you have the `wikipedia_scraper.py` file in your working directory.

2. **Install Dependencies**

   Install the required Python libraries:

   ```bash
   pip install requests beautifulsoup4
   ```

3. **Run the Script**

   Execute the script using Python:

   ```bash
   python wikipedia_scraper.py
   ```

4. **Output**

   The extracted articles will be saved in an `articles.jsonl` file in the same directory.

---

### Output Format

The `articles.jsonl` file contains the extracted text in JSON Lines format. Each line is a JSON object with the following structure:

```json
{
  "url": "<URL of the Wikipedia page>",
  "text": "<Extracted text content>"
}
```

---

### Sample Entry

```json
{
  "url": "https://en.wikipedia.org/wiki/Robotics",
  "text": "Robotics is an interdisciplinary branch of computer science and engineering. Robotics involves the design, construction, operation, and use of robots..."
}
```

---

### References

- **Go and Colly**:

  - [Go Programming Language](https://golang.org/)
  - [Colly Web Scraping Framework](http://go-colly.org/docs/)

- **Python Libraries**:

  - [Requests Library Documentation](https://docs.python-requests.org/en/master/)
  - [BeautifulSoup Documentation](https://www.crummy.com/software/BeautifulSoup/bs4/doc/)

- **Web Scraping Best Practices**:

  - [Scraping and Crawling Guidelines](https://www.scrapinghub.com/what-is-web-scraping)

---