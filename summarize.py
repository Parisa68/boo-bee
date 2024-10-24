import logging
import os
import sys
from transformers import pipeline


os.environ['TF_CPP_MIN_LOG_LEVEL'] = '3'
logging.getLogger().setLevel(logging.ERROR)
logging.getLogger("transformers").setLevel(logging.ERROR)
logging.getLogger("tensorflow").setLevel(logging.ERROR)

stderr = sys.stderr
with open(os.devnull, 'w') as devnull:
    sys.stderr = devnull

    def summarize(article):
        try:
            summarizer = pipeline("summarization", model="facebook/bart-large-cnn")
            summary = summarizer(article, min_length=300, max_length=400, do_sample=False)
            return summary[0]['summary_text']
        except Exception as e:
            return f"Error during summarization: {str(e)}"

    if __name__ == "__main__":
        if len(sys.argv) < 2:
            print("Please provide text to summarize")
            sys.exit(1)

        article = sys.argv[1]
        summary = summarize(article)
        print(summary.strip(), end='')

sys.stderr = stderr