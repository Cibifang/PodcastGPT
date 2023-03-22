import unittest
from rss_reader import get_list_from_rss, download_selected_podcasts_from_rss

class TestRSSReader(unittest.TestCase):

    def test_get_list_from_rss(self):
        rss_url = "https://feeds.fireside.fm/stovol/rss"
        podcast_list = get_list_from_rss(rss_url)
        print(podcast_list)
        self.assertIsInstance(podcast_list, list)
        self.assertTrue(all(isinstance(p, str) for p in podcast_list))

    def test_download_selected_podcasts_from_rss(self):
        rss_url = "https://feeds.fireside.fm/stovol/rss"
        selected_files = ["【随机波动】回到街道间，回看好莱坞：时尚、电影与性别", "【随机波动021】桌上一堆书，对影成三人"]
        result = download_selected_podcasts_from_rss(rss_url, selected_files, "test_podcast/")
        self.assertTrue(result)

if __name__ == '__main__':
    unittest.main()
