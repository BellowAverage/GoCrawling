import requests
from bs4 import BeautifulSoup
import json
import time

# List of Wikipedia URLs for topics of interest
urls = [
    "https://en.wikipedia.org/wiki/Robotics",
    "https://en.wikipedia.org/wiki/Robot",
    "https://en.wikipedia.org/wiki/Reinforcement_learning",
    "https://en.wikipedia.org/wiki/Robot_Operating_System",
    "https://en.wikipedia.org/wiki/Intelligent_agent",
    "https://en.wikipedia.org/wiki/Software_agent",
    "https://en.wikipedia.org/wiki/Robotic_process_automation",
    "https://en.wikipedia.org/wiki/Chatbot",
    "https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
    "https://en.wikipedia.org/wiki/Android_(robot)"
]

def scrape_wikipedia_articles(urls):
    start_time = time.time()
    articles = []

    for url in urls:
        print(f"Processing {url}")
        try:
            # Fetch the page content
            response = requests.get(url)
            response.raise_for_status()  # Check for HTTP errors

            # Parse the HTML content using BeautifulSoup
            soup = BeautifulSoup(response.content, 'html.parser')

            # Find the main content div
            content_div = soup.find('div', {'id': 'mw-content-text'})

            # Extract text from paragraphs and list items
            paragraphs = content_div.find_all(['p', 'li'])
            text = ' '.join([para.get_text(strip=True) for para in paragraphs])

            # Create a dictionary for the article
            article = {
                'url': url,
                'text': text
            }

            articles.append(article)

        except Exception as e:
            print(f"An error occurred while processing {url}: {e}")

    # Write the articles to a JSON lines file
    with open('articles_python.jsonl', 'w', encoding='utf-8') as f:
        for article in articles:
            json_line = json.dumps(article, ensure_ascii=False)
            f.write(json_line + '\n')

    elapsed_time = time.time() - start_time
    print(f"Execution time: {elapsed_time:.2f} seconds")

if __name__ == "__main__":
    scrape_wikipedia_articles(urls)
