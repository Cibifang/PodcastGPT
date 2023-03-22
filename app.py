"""
This file defines a Flask application with three routes:

/: The root route that renders the index.html template.
/get_rss_list: A route that get file list from the specified RSS feed.
/download_podcasts: A route that downloads selected podcast content from the specified RSS feed.
/summarize_podcast: A route that summarizes the podcast content.
/text_to_speech: A route that converts text to speech.
The download_podcasts route takes a POST request with a form parameter rss_url that specifies the URL of the RSS feed to download. The download_podcasts_from_rss function is called to download the podcast content from the specified RSS feed.

The summarize_podcast route takes a POST request with a file parameter audio_file that contains the podcast content. The recognize_speech function is called to convert the audio file to text, and the summarize_text function is called to summarize the text.

The text_to_speech route takes a POST request with a form parameter text that contains the text to convert to speech. The text_to_speech function is called to convert the text to speech.  
"""

from flask import Flask, render_template, request
from rss_reader import get_list_from_rss, download_selected_podcasts_from_rss
from podcast_recognition import recognize_speech
from text_to_speech import text_to_speech
from api_call import summarize_text

app = Flask(__name__)

@app.route("/")
def index():
    return render_template("index.html")

@app.route("/get_rss_list", methods=["POST"])
def get_rss_list():
    rss_url = request.form["rss_url"]
    podcast_list = get_list_from_rss(rss_url)
    return jsonify(podcast_list)

@app.route("/download_selected_podcasts", methods=["POST"])
def download_selected_podcasts():
    selected_files = request.json.get("selected_files")
    download_selected_podcasts_from_rss(rss_url, selected_files, "podcast/")
    return "Podcasts downloaded successfully!"

@app.route("/summarize_podcast", methods=["POST"])
def summarize_podcast():
    audio_file = request.files["audio_file"]
    text = recognize_speech(audio_file)
    summary = summarize_text(text)
    return summary

@app.route("/text_to_speech", methods=["POST"])
def convert_text_to_speech():
    text = request.form["text"]
    audio_file = text_to_speech(text)
    return audio_file

if __name__ == "__main__":
    app.run(debug=True)
