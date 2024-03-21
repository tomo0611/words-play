import os

import json

f = open('words.json', 'r', encoding='utf-8')
data = json.load(f)
f.close()

counter = 0

word = ""
word_ja = ""
audio = ""

out_json = {
    "words": []
}

for word in data:
    # module_type': 'reversibleCard
    # module_type': 'audio
    if data[word]['module_type'] == 'reversibleCard':
        word_ja = (data[word]['contents'][0]['contents'][0]['contents'].replace('<div style="font-size:1.48em;text-align:center;color:#1b0297;font-weight:bold;"><p>',"").replace("</p></div>\n",""))
        word_en = (data[word]['contents'][1]['contents'][0]['contents'].replace('<div style="font-size:1.14em;text-align:center;color:#133922;font-weight:normal;"><p>',"").replace("</p></div>\n",""))
    elif data[word]['module_type'] == 'audio':
        audio = "https://singlepacker.com/p/5gr0/obunshadigital/m826"+(data[word]['contents'])[1:]
        out_json["words"].append(
            {
                "word_ja": word_ja,
                "word_en": word_en,
                "audio_url": audio
            }
        )
        counter += 1

print(counter)

with open('words_out.json', 'w', encoding='utf-8') as f:
    json.dump(out_json, f, ensure_ascii=False, indent=4)
    f.close()
    