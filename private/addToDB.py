import json
import mysql.connector

f = open("words_out.json", "r", encoding="utf-8")
data = json.load(f)
f.close()

conn = mysql.connector.connect(
    host="localhost",
    user="wordsplay",
    password="password"
)
cursor = conn.cursor()
cursor.execute("USE WORDS_PLAY;")

id = 1
for word in data["words"]:
    cursor.execute("INSERT INTO words (id, word_en, word_ja, audio_url) VALUES (%s,%s, %s,%s);", (id, word["word_en"], word["word_ja"], word["audio_url"]))
    id=id+1

conn.commit()
cursor.close()
conn.close()