# Implement audio format conversion

import os
from pydub import AudioSegment

def convert_audio_to_wav(audio_file):
    """
    Convert an audio file to WAV format.
    :param audio_file: str, the path of the audio file to convert
    :return: str, the path of the converted audio file
    """
    # get the file name and extension
    file_name, file_extension = os.path.splitext(audio_file)

    # load the audio file
    sound = AudioSegment.from_file(audio_file, format=file_extension[1:])

    # set the sample rate to 16000 Hz and convert the audio file to 16-bit mono
    sound = sound.set_frame_rate(16000).set_channels(1)

    # save the audio file in WAV format
    wav_file = file_name + '.wav'
    sound.export(wav_file, format='wav')

    return wav_file
