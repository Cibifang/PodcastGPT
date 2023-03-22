import unittest
from podcast_recognition import recognize_speech

class TestSpeechRecognition(unittest.TestCase):
    def test_recognize_speech(self):
        audio_file = "test_podcast/【随机波动】回到街道间，回看好莱坞：时尚、电影与性别.mp3"
        expected_text = "this is a test"
        recognized_text = recognize_speech(audio_file)
        print(recognized_text)
        self.assertEqual(recognized_text, expected_text)

if __name__ == "__main__":
    unittest.main()