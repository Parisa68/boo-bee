from transformers import pipeline

def summarize(article):
    summarizer = pipeline("summarization", model="facebook/bart-large-cnn")
    summary = summarizer(article, min_length=300, do_sample=False)
    return summary[0]['summary_text']

if __name__ == "__main__":
    import sys
    article = sys.argv[1]
    print(summarize(article))