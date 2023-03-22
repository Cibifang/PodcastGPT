"""
Get podcast content from the specified RSS and download podcast content
"""

import feedparser
import os
import requests

def get_list_from_rss(rss_url):
    """
    Get the list of podcasts from an RSS feed and return their titles.
    :param rss_url: str, the URL of the RSS feed
    :return: list, the list of podcast titles
    """
    podcast_list = []
    feed = feedparser.parse(rss_url)
    for entry in feed.entries:
        title = entry.title
        podcast_list.append(title)
    return podcast_list


def download_selected_podcasts_from_rss(rss_url, selected_files, save_dir):
    """
    Download the selected podcasts from an RSS feed and save them to the specified directory.
    :param rss_url: str, the URL of the RSS feed
    :param selected_files: list, the list of selected podcast titles to download
    :param save_dir: str, the directory to save the downloaded podcasts
    :return: None
    """
    feed = feedparser.parse(rss_url)
    print("selected: ", selected_files)
    for entry in feed.entries:
        title = entry.title
        url = entry.enclosures[0].href
        if title in selected_files:
            print(title)
            file_path = os.path.join(save_dir, title + os.path.splitext(url)[1])
            with open(file_path, "wb") as f:
                f.write(requests.get(url).content)