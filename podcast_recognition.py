"""
Implement speech recognition content and convert it to text.
"""

import speech_recognition as sr
from audio_converter import convert_audio_to_wav

def recognize_speech(audio_file):
    """
    Recognize the speech in an audio file.
    :param audio_file: str, the path of the audio file to recognize
    :return: str, the recognized text
    """
    # convert the audio file to WAV format
    wav_file = convert_audio_to_wav(audio_file)

    # recognize the speech in the WAV file
    r = sr.Recognizer()
    with sr.AudioFile(wav_file) as source:
        audio = r.record(source)
    return r.recognize_sphinx(audio, language='zh-CN')
