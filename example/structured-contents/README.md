# Structured-contents API examples
Used to fetch pre-parsed articles in their latest revision/version from all supported projects and languages. 
Allows filtering and field selection. Allows to limit articles when doing cross-project, cross-language lookup.
Refer to the documentation [here](https://enterprise.wikimedia.com/docs/on-demand/#article-structured-contents-beta).
The articles are based on [this](https://github.com/wikimedia-enterprise/wikimedia-enterprise/blob/main/general/schema/article.go) schema.



i) Get pre-parsed articles with name `Montreal` from all projects in English language. 

```bash
POST https://api.enterprise.wikimedia.com/v2/structured-contents/Montreal
```

with request parameter:
```json
{
    "limit": 5,
    "filters": [
        {
            "field": "in_language.identifier",
            "value": "en"
        }
    ]
}
```
Note: Max allowed limit is 10. Default limit is 3.

<detail>
<summary>Response:</summary> 

```json

[
    {
        "name": "Montreal",
        "identifier": 4344,
        "date_created": "2003-04-21T16:47:27Z",
        "date_modified": "2024-02-29T04:54:52Z",
        "date_previously_modified": "2024-01-31T11:47:59Z",
        "version": {
            "identifier": 78276816,
            "comment": "/* Translations */+gujarati, based on https://helloenglish.com/translate/english-gujarati-dictionary/meaning-of-montreal-in-Gujarati/amp and also some limited familiarity with the writing system",
            "tags": [
                "mobile edit",
                "mobile web edit"
            ],
            "scores": {},
            "editor": {
                "identifier": 3026,
                "name": "Jun-Dai",
                "edit_count": 1552,
                "groups": [
                    "autopatrolled",
                    "*",
                    "user",
                    "autoconfirmed"
                ],
                "date_started": "0001-01-01T00:00:00Z"
            },
            "number_of_characters": 8884,
            "size": {
                "value": 9267,
                "unit_text": "B"
            }
        },
        "previous_version": {
            "identifier": 77790591,
            "number_of_characters": 8852
        },
        "url": "https://en.wiktionary.org/wiki/Montreal",
        "namespace": {
            "identifier": 0
        },
        "in_language": {
            "identifier": "en"
        },
        "additional_entities": [
            {
                "identifier": "Q652",
                "url": "https://www.wikidata.org/entity/Q652",
                "aspects": [
                    "S"
                ]
            },
            {
                "identifier": "Q1473289",
                "url": "https://www.wikidata.org/entity/Q1473289",
                "aspects": [
                    "S"
                ]
            },
            {
                "identifier": "Q150",
                "url": "https://www.wikidata.org/entity/Q150",
                "aspects": [
                    "S"
                ]
            },
            {
                "identifier": "Q1860",
                "url": "https://www.wikidata.org/entity/Q1860",
                "aspects": [
                    "S"
                ]
            }
        ],
        "categories": [
            {
                "name": "Category:Polish links with redundant wikilinks",
                "url": "https://en.wiktionary.org/wiki/Category:Polish_links_with_redundant_wikilinks"
            },
            {
                "name": "Category:Portuguese terms borrowed from French",
                "url": "https://en.wiktionary.org/wiki/Category:Portuguese_terms_borrowed_from_French"
            },
            {
                "name": "Category:Rhymes:English/ɔːl/3 syllables",
                "url": "https://en.wiktionary.org/wiki/Category:Rhymes:English/ɔːl/3_syllables"
            },
            {
                "name": "Category:en:Rivers in the United States",
                "url": "https://en.wiktionary.org/wiki/Category:en:Rivers_in_the_United_States"
            },
            {
                "name": "Category:pl:Montreal",
                "url": "https://en.wiktionary.org/wiki/Category:pl:Montreal"
            },
            {
                "name": "Category:Rhymes:Polish/ɛal",
                "url": "https://en.wiktionary.org/wiki/Category:Rhymes:Polish/ɛal"
            },
            {
                "name": "Category:es:Places in Canada",
                "url": "https://en.wiktionary.org/wiki/Category:es:Places_in_Canada"
            },
            {
                "name": "Category:Forms linking to themselves",
                "url": "https://en.wiktionary.org/wiki/Category:Forms_linking_to_themselves"
            },
            {
                "name": "Category:Turkish terms with IPA pronunciation",
                "url": "https://en.wiktionary.org/wiki/Category:Turkish_terms_with_IPA_pronunciation"
            },
            {
                "name": "Category:de:Cities in Quebec",
                "url": "https://en.wiktionary.org/wiki/Category:de:Cities_in_Quebec"
            },
            {
                "name": "Category:de:Cities in Canada",
                "url": "https://en.wiktionary.org/wiki/Category:de:Cities_in_Canada"
            },
            {
                "name": "Category:Polish terms borrowed from English",
                "url": "https://en.wiktionary.org/wiki/Category:Polish_terms_borrowed_from_English"
            },
            {
                "name": "Category:en:Cities in the United States",
                "url": "https://en.wiktionary.org/wiki/Category:en:Cities_in_the_United_States"
            },
            {
                "name": "Category:Catalan proper nouns",
                "url": "https://en.wiktionary.org/wiki/Category:Catalan_proper_nouns"
            },
            {
                "name": "Category:ca:Islands",
                "url": "https://en.wiktionary.org/wiki/Category:ca:Islands"
            },
            {
                "name": "Category:English 3-syllable words",
                "url": "https://en.wiktionary.org/wiki/Category:English_3-syllable_words"
            },
            {
                "name": "Category:Polish lemmas",
                "url": "https://en.wiktionary.org/wiki/Category:Polish_lemmas"
            },
            {
                "name": "Category:Portuguese 4-syllable words",
                "url": "https://en.wiktionary.org/wiki/Category:Portuguese_4-syllable_words"
            },
            {
                "name": "Category:pt:Places in Quebec",
                "url": "https://en.wiktionary.org/wiki/Category:pt:Places_in_Quebec"
            },
            {
                "name": "Category:en:Montreal",
                "url": "https://en.wiktionary.org/wiki/Category:en:Montreal"
            },
            {
                "name": "Category:Middle French lemmas",
                "url": "https://en.wiktionary.org/wiki/Category:Middle_French_lemmas"
            },
            {
                "name": "Category:pl:Places in Quebec",
                "url": "https://en.wiktionary.org/wiki/Category:pl:Places_in_Quebec"
            },
            {
                "name": "Category:en:Places in Quebec",
                "url": "https://en.wiktionary.org/wiki/Category:en:Places_in_Quebec"
            },
            {
                "name": "Category:Polish links with manual fragments",
                "url": "https://en.wiktionary.org/wiki/Category:Polish_links_with_manual_fragments"
            },
            {
                "name": "Category:Slovak lemmas",
                "url": "https://en.wiktionary.org/wiki/Category:Slovak_lemmas"
            },
            {
                "name": "Category:pt:Cities in Canada",
                "url": "https://en.wiktionary.org/wiki/Category:pt:Cities_in_Canada"
            },
            {
                "name": "Category:en:Rivers in Saskatchewan",
                "url": "https://en.wiktionary.org/wiki/Category:en:Rivers_in_Saskatchewan"
            },
            {
                "name": "Category:nb:Places in Quebec",
                "url": "https://en.wiktionary.org/wiki/Category:nb:Places_in_Quebec"
            },
            {
                "name": "Category:pt:Cities in Quebec",
                "url": "https://en.wiktionary.org/wiki/Category:pt:Cities_in_Quebec"
            },
            {
                "name": "Category:ca:Cities in Quebec",
                "url": "https://en.wiktionary.org/wiki/Category:ca:Cities_in_Quebec"
            },
            {
                "name": "Category:ca:Cities in Canada",
                "url": "https://en.wiktionary.org/wiki/Category:ca:Cities_in_Canada"
            },
            {
                "name": "Category:en:Cities in Quebec",
                "url": "https://en.wiktionary.org/wiki/Category:en:Cities_in_Quebec"
            },
            {
                "name": "Category:en:Cities in Canada",
                "url": "https://en.wiktionary.org/wiki/Category:en:Cities_in_Canada"
            },
            {
                "name": "Category:Arabic terms with non-redundant manual transliterations",
                "url": "https://en.wiktionary.org/wiki/Category:Arabic_terms_with_non-redundant_manual_transliterations"
            },
            {
                "name": "Category:Polish 3-syllable words",
                "url": "https://en.wiktionary.org/wiki/Category:Polish_3-syllable_words"
            },
            {
                "name": "Category:sk:Places in Quebec",
                "url": "https://en.wiktionary.org/wiki/Category:sk:Places_in_Quebec"
            },
            {
                "name": "Category:en:Places in the United States",
                "url": "https://en.wiktionary.org/wiki/Category:en:Places_in_the_United_States"
            },
            {
                "name": "Category:ca:Places in Quebec",
                "url": "https://en.wiktionary.org/wiki/Category:ca:Places_in_Quebec"
            },
            {
                "name": "Category:German terms with IPA pronunciation",
                "url": "https://en.wiktionary.org/wiki/Category:German_terms_with_IPA_pronunciation"
            },
            {
                "name": "Category:nn:Cities in Quebec",
                "url": "https://en.wiktionary.org/wiki/Category:nn:Cities_in_Quebec"
            },
            {
                "name": "Category:nn:Cities in Canada",
                "url": "https://en.wiktionary.org/wiki/Category:nn:Cities_in_Canada"
            },
            {
                "name": "Category:Spanish proper nouns",
                "url": "https://en.wiktionary.org/wiki/Category:Spanish_proper_nouns"
            },
            {
                "name": "Category:English terms derived from Middle French",
                "url": "https://en.wiktionary.org/wiki/Category:English_terms_derived_from_Middle_French"
            },
            {
                "name": "Category:English proper nouns",
                "url": "https://en.wiktionary.org/wiki/Category:English_proper_nouns"
            },
            {
                "name": "Category:Norwegian Nynorsk proper nouns",
                "url": "https://en.wiktionary.org/wiki/Category:Norwegian_Nynorsk_proper_nouns"
            },
            {
                "name": "Category:Turkish lemmas",
                "url": "https://en.wiktionary.org/wiki/Category:Turkish_lemmas"
            },
            {
                "name": "Category:en:Unincorporated communities in the United States",
                "url": "https://en.wiktionary.org/wiki/Category:en:Unincorporated_communities_in_the_United_States"
            },
            {
                "name": "Category:nn:Places in Quebec",
                "url": "https://en.wiktionary.org/wiki/Category:nn:Places_in_Quebec"
            },
            {
                "name": "Category:Polish links with redundant alt parameters",
                "url": "https://en.wiktionary.org/wiki/Category:Polish_links_with_redundant_alt_parameters"
            },
            {
                "name": "Category:frm:Islands",
                "url": "https://en.wiktionary.org/wiki/Category:frm:Islands"
            },
            {
                "name": "Category:en:Unincorporated communities in Missouri, USA",
                "url": "https://en.wiktionary.org/wiki/Category:en:Unincorporated_communities_in_Missouri,_USA"
            },
            {
                "name": "Category:Mandarin terms with redundant transliterations",
                "url": "https://en.wiktionary.org/wiki/Category:Mandarin_terms_with_redundant_transliterations"
            },
            {
                "name": "Category:Middle French compound terms",
                "url": "https://en.wiktionary.org/wiki/Category:Middle_French_compound_terms"
            },
            {
                "name": "Category:Portuguese terms with IPA pronunciation",
                "url": "https://en.wiktionary.org/wiki/Category:Portuguese_terms_with_IPA_pronunciation"
            },
            {
                "name": "Category:Slovak masculine nouns",
                "url": "https://en.wiktionary.org/wiki/Category:Slovak_masculine_nouns"
            },
            {
                "name": "Category:es:Cities in Quebec",
                "url": "https://en.wiktionary.org/wiki/Category:es:Cities_in_Quebec"
            },
            {
                "name": "Category:es:Cities in Canada",
                "url": "https://en.wiktionary.org/wiki/Category:es:Cities_in_Canada"
            },
            {
                "name": "Category:en:Rivers in Wisconsin, USA",
                "url": "https://en.wiktionary.org/wiki/Category:en:Rivers_in_Wisconsin,_USA"
            },
            {
                "name": "Category:Polish terms derived from English",
                "url": "https://en.wiktionary.org/wiki/Category:Polish_terms_derived_from_English"
            },
            {
                "name": "Category:Polish terms with audio links",
                "url": "https://en.wiktionary.org/wiki/Category:Polish_terms_with_audio_links"
            },
            {
                "name": "Category:ca:Montreal",
                "url": "https://en.wiktionary.org/wiki/Category:ca:Montreal"
            },
            {
                "name": "Category:German terms with audio links",
                "url": "https://en.wiktionary.org/wiki/Category:German_terms_with_audio_links"
            },
            {
                "name": "Category:Rhymes:Spanish/al/3 syllables",
                "url": "https://en.wiktionary.org/wiki/Category:Rhymes:Spanish/al/3_syllables"
            },
            {
                "name": "Category:nb:Montreal",
                "url": "https://en.wiktionary.org/wiki/Category:nb:Montreal"
            },
            {
                "name": "Category:Polish masculine nouns",
                "url": "https://en.wiktionary.org/wiki/Category:Polish_masculine_nouns"
            },
            {
                "name": "Category:Slovak terms with IPA pronunciation",
                "url": "https://en.wiktionary.org/wiki/Category:Slovak_terms_with_IPA_pronunciation"
            },
            {
                "name": "Category:Rhymes:Spanish/al",
                "url": "https://en.wiktionary.org/wiki/Category:Rhymes:Spanish/al"
            },
            {
                "name": "Category:en:Rivers in Michigan, USA",
                "url": "https://en.wiktionary.org/wiki/Category:en:Rivers_in_Michigan,_USA"
            },
            {
                "name": "Category:de:Montreal",
                "url": "https://en.wiktionary.org/wiki/Category:de:Montreal"
            },
            {
                "name": "Category:de:Places in Canada",
                "url": "https://en.wiktionary.org/wiki/Category:de:Places_in_Canada"
            },
            {
                "name": "Category:Rhymes:Polish/ɛal/3 syllables",
                "url": "https://en.wiktionary.org/wiki/Category:Rhymes:Polish/ɛal/3_syllables"
            },
            {
                "name": "Category:Polish inanimate nouns",
                "url": "https://en.wiktionary.org/wiki/Category:Polish_inanimate_nouns"
            },
            {
                "name": "Category:Portuguese feminine nouns",
                "url": "https://en.wiktionary.org/wiki/Category:Portuguese_feminine_nouns"
            },
            {
                "name": "Category:Turkish nouns",
                "url": "https://en.wiktionary.org/wiki/Category:Turkish_nouns"
            },
            {
                "name": "Category:English terms with audio links",
                "url": "https://en.wiktionary.org/wiki/Category:English_terms_with_audio_links"
            },
            {
                "name": "Category:German lemmas",
                "url": "https://en.wiktionary.org/wiki/Category:German_lemmas"
            },
            {
                "name": "Category:Polish terms derived from French",
                "url": "https://en.wiktionary.org/wiki/Category:Polish_terms_derived_from_French"
            },
            {
                "name": "Category:Slovak proper nouns",
                "url": "https://en.wiktionary.org/wiki/Category:Slovak_proper_nouns"
            },
            {
                "name": "Category:en:Islands",
                "url": "https://en.wiktionary.org/wiki/Category:en:Islands"
            },
            {
                "name": "Category:frm:Places in Canada",
                "url": "https://en.wiktionary.org/wiki/Category:frm:Places_in_Canada"
            },
            {
                "name": "Category:Polish terms with IPA pronunciation",
                "url": "https://en.wiktionary.org/wiki/Category:Polish_terms_with_IPA_pronunciation"
            },
            {
                "name": "Category:nb:Cities in Canada",
                "url": "https://en.wiktionary.org/wiki/Category:nb:Cities_in_Canada"
            },
            {
                "name": "Category:Slovak inanimate nouns",
                "url": "https://en.wiktionary.org/wiki/Category:Slovak_inanimate_nouns"
            },
            {
                "name": "Category:es:Places in Quebec",
                "url": "https://en.wiktionary.org/wiki/Category:es:Places_in_Quebec"
            },
            {
                "name": "Category:Inuktitut terms with redundant script codes",
                "url": "https://en.wiktionary.org/wiki/Category:Inuktitut_terms_with_redundant_script_codes"
            },
            {
                "name": "Category:Norwegian Bokmål lemmas",
                "url": "https://en.wiktionary.org/wiki/Category:Norwegian_Bokmål_lemmas"
            },
            {
                "name": "Category:nb:Cities in Quebec",
                "url": "https://en.wiktionary.org/wiki/Category:nb:Cities_in_Quebec"
            },
            {
                "name": "Category:English terms with historical senses",
                "url": "https://en.wiktionary.org/wiki/Category:English_terms_with_historical_senses"
            },
            {
                "name": "Category:Portuguese proper nouns",
                "url": "https://en.wiktionary.org/wiki/Category:Portuguese_proper_nouns"
            },
            {
                "name": "Category:Catalan terms with IPA pronunciation",
                "url": "https://en.wiktionary.org/wiki/Category:Catalan_terms_with_IPA_pronunciation"
            },
            {
                "name": "Category:nb:Places in Canada",
                "url": "https://en.wiktionary.org/wiki/Category:nb:Places_in_Canada"
            },
            {
                "name": "Category:Norwegian Nynorsk lemmas",
                "url": "https://en.wiktionary.org/wiki/Category:Norwegian_Nynorsk_lemmas"
            },
            {
                "name": "Category:Polish singularia tantum",
                "url": "https://en.wiktionary.org/wiki/Category:Polish_singularia_tantum"
            },
            {
                "name": "Category:Portuguese 2-syllable words",
                "url": "https://en.wiktionary.org/wiki/Category:Portuguese_2-syllable_words"
            },
            {
                "name": "Category:es:Montreal",
                "url": "https://en.wiktionary.org/wiki/Category:es:Montreal"
            },
            {
                "name": "Category:Turkish terms derived from French",
                "url": "https://en.wiktionary.org/wiki/Category:Turkish_terms_derived_from_French"
            },
            {
                "name": "Category:en:Places in Canada",
                "url": "https://en.wiktionary.org/wiki/Category:en:Places_in_Canada"
            },
            {
                "name": "Category:en:Rivers in Ontario",
                "url": "https://en.wiktionary.org/wiki/Category:en:Rivers_in_Ontario"
            },
            {
                "name": "Category:German proper nouns",
                "url": "https://en.wiktionary.org/wiki/Category:German_proper_nouns"
            },
            {
                "name": "Category:pl:Places in Canada",
                "url": "https://en.wiktionary.org/wiki/Category:pl:Places_in_Canada"
            },
            {
                "name": "Category:pt:Places in Canada",
                "url": "https://en.wiktionary.org/wiki/Category:pt:Places_in_Canada"
            },
            {
                "name": "Category:Spanish lemmas",
                "url": "https://en.wiktionary.org/wiki/Category:Spanish_lemmas"
            },
            {
                "name": "Category:en:Cities in Wisconsin, USA",
                "url": "https://en.wiktionary.org/wiki/Category:en:Cities_in_Wisconsin,_USA"
            },
            {
                "name": "Category:en:Places in Wisconsin, USA",
                "url": "https://en.wiktionary.org/wiki/Category:en:Places_in_Wisconsin,_USA"
            },
            {
                "name": "Category:nn:Montreal",
                "url": "https://en.wiktionary.org/wiki/Category:nn:Montreal"
            },
            {
                "name": "Category:en:Rivers in Canada",
                "url": "https://en.wiktionary.org/wiki/Category:en:Rivers_in_Canada"
            },
            {
                "name": "Category:Cantonese terms with redundant transliterations",
                "url": "https://en.wiktionary.org/wiki/Category:Cantonese_terms_with_redundant_transliterations"
            },
            {
                "name": "Category:pt:Montreal",
                "url": "https://en.wiktionary.org/wiki/Category:pt:Montreal"
            },
            {
                "name": "Category:en:Places in Ontario",
                "url": "https://en.wiktionary.org/wiki/Category:en:Places_in_Ontario"
            },
            {
                "name": "Category:Norwegian Bokmål proper nouns",
                "url": "https://en.wiktionary.org/wiki/Category:Norwegian_Bokmål_proper_nouns"
            },
            {
                "name": "Category:nn:Places in Canada",
                "url": "https://en.wiktionary.org/wiki/Category:nn:Places_in_Canada"
            },
            {
                "name": "Category:en:Historical political subdivisions",
                "url": "https://en.wiktionary.org/wiki/Category:en:Historical_political_subdivisions"
            },
            {
                "name": "Category:Portuguese 3-syllable words",
                "url": "https://en.wiktionary.org/wiki/Category:Portuguese_3-syllable_words"
            },
            {
                "name": "Category:ca:Places in Canada",
                "url": "https://en.wiktionary.org/wiki/Category:ca:Places_in_Canada"
            },
            {
                "name": "Category:Portuguese terms derived from French",
                "url": "https://en.wiktionary.org/wiki/Category:Portuguese_terms_derived_from_French"
            },
            {
                "name": "Category:sk:Places in Canada",
                "url": "https://en.wiktionary.org/wiki/Category:sk:Places_in_Canada"
            },
            {
                "name": "Category:English terms with IPA pronunciation",
                "url": "https://en.wiktionary.org/wiki/Category:English_terms_with_IPA_pronunciation"
            },
            {
                "name": "Category:English lemmas",
                "url": "https://en.wiktionary.org/wiki/Category:English_lemmas"
            },
            {
                "name": "Category:English uncountable nouns",
                "url": "https://en.wiktionary.org/wiki/Category:English_uncountable_nouns"
            },
            {
                "name": "Category:German neuter nouns",
                "url": "https://en.wiktionary.org/wiki/Category:German_neuter_nouns"
            },
            {
                "name": "Category:Middle French proper nouns",
                "url": "https://en.wiktionary.org/wiki/Category:Middle_French_proper_nouns"
            },
            {
                "name": "Category:Polish proper nouns",
                "url": "https://en.wiktionary.org/wiki/Category:Polish_proper_nouns"
            },
            {
                "name": "Category:sk:Cities in Quebec",
                "url": "https://en.wiktionary.org/wiki/Category:sk:Cities_in_Quebec"
            },
            {
                "name": "Category:sk:Cities in Canada",
                "url": "https://en.wiktionary.org/wiki/Category:sk:Cities_in_Canada"
            },
            {
                "name": "Category:English terms derived from French",
                "url": "https://en.wiktionary.org/wiki/Category:English_terms_derived_from_French"
            },
            {
                "name": "Category:English terms derived from Italian",
                "url": "https://en.wiktionary.org/wiki/Category:English_terms_derived_from_Italian"
            },
            {
                "name": "Category:Catalan lemmas",
                "url": "https://en.wiktionary.org/wiki/Category:Catalan_lemmas"
            },
            {
                "name": "Category:Spanish terms with IPA pronunciation",
                "url": "https://en.wiktionary.org/wiki/Category:Spanish_terms_with_IPA_pronunciation"
            },
            {
                "name": "Category:Spanish feminine nouns",
                "url": "https://en.wiktionary.org/wiki/Category:Spanish_feminine_nouns"
            },
            {
                "name": "Category:en:Places in Saskatchewan",
                "url": "https://en.wiktionary.org/wiki/Category:en:Places_in_Saskatchewan"
            },
            {
                "name": "Category:de:Places in Quebec",
                "url": "https://en.wiktionary.org/wiki/Category:de:Places_in_Quebec"
            },
            {
                "name": "Category:Spanish 3-syllable words",
                "url": "https://en.wiktionary.org/wiki/Category:Spanish_3-syllable_words"
            },
            {
                "name": "Category:Portuguese lemmas",
                "url": "https://en.wiktionary.org/wiki/Category:Portuguese_lemmas"
            },
            {
                "name": "Category:sk:Montreal",
                "url": "https://en.wiktionary.org/wiki/Category:sk:Montreal"
            },
            {
                "name": "Category:English terms borrowed from French",
                "url": "https://en.wiktionary.org/wiki/Category:English_terms_borrowed_from_French"
            },
            {
                "name": "Category:English ellipses",
                "url": "https://en.wiktionary.org/wiki/Category:English_ellipses"
            },
            {
                "name": "Category:Catalan masculine nouns",
                "url": "https://en.wiktionary.org/wiki/Category:Catalan_masculine_nouns"
            },
            {
                "name": "Category:pl:Cities in Canada",
                "url": "https://en.wiktionary.org/wiki/Category:pl:Cities_in_Canada"
            },
            {
                "name": "Category:en:Places in Missouri, USA",
                "url": "https://en.wiktionary.org/wiki/Category:en:Places_in_Missouri,_USA"
            },
            {
                "name": "Category:en:Places in Michigan, USA",
                "url": "https://en.wiktionary.org/wiki/Category:en:Places_in_Michigan,_USA"
            },
            {
                "name": "Category:pl:Cities in Quebec",
                "url": "https://en.wiktionary.org/wiki/Category:pl:Cities_in_Quebec"
            },
            {
                "name": "Category:Rhymes:English/ɔːl",
                "url": "https://en.wiktionary.org/wiki/Category:Rhymes:English/ɔːl"
            }
        ],
        "templates": [
            {
                "name": "Template:R:sk:SDK",
                "url": "https://en.wiktionary.org/wiki/Template:R:sk:SDK"
            },
            {
                "name": "Template:wikipedia",
                "url": "https://en.wiktionary.org/wiki/Template:wikipedia"
            },
            {
                "name": "Template:cot",
                "url": "https://en.wiktionary.org/wiki/Template:cot"
            },
            {
                "name": "Template:sk-proper noun",
                "url": "https://en.wiktionary.org/wiki/Template:sk-proper_noun"
            },
            {
                "name": "Template:alt form",
                "url": "https://en.wiktionary.org/wiki/Template:alt_form"
            },
            {
                "name": "Template:bor ",
                "url": "https://en.wiktionary.org/wiki/Template:bor+"
            },
            {
                "name": "Template:alt",
                "url": "https://en.wiktionary.org/wiki/Template:alt"
            },
            {
                "name": "Template:alter",
                "url": "https://en.wiktionary.org/wiki/Template:alter"
            },
            {
                "name": "Template:IPA",
                "url": "https://en.wiktionary.org/wiki/Template:IPA"
            },
            {
                "name": "Template:der-bottom",
                "url": "https://en.wiktionary.org/wiki/Template:der-bottom"
            },
            {
                "name": "Template:es-proper noun",
                "url": "https://en.wiktionary.org/wiki/Template:es-proper_noun"
            },
            {
                "name": "Template:syn",
                "url": "https://en.wiktionary.org/wiki/Template:syn"
            },
            {
                "name": "Template:de-proper noun",
                "url": "https://en.wiktionary.org/wiki/Template:de-proper_noun"
            },
            {
                "name": "Template:pl-decl-noun-m-in",
                "url": "https://en.wiktionary.org/wiki/Template:pl-decl-noun-m-in"
            },
            {
                "name": "Template:t ",
                "url": "https://en.wiktionary.org/wiki/Template:t+"
            },
            {
                "name": "Template:bottom",
                "url": "https://en.wiktionary.org/wiki/Template:bottom"
            },
            {
                "name": "Template:pl-p",
                "url": "https://en.wiktionary.org/wiki/Template:pl-p"
            },
            {
                "name": "Template:also",
                "url": "https://en.wiktionary.org/wiki/Template:also"
            },
            {
                "name": "Template:lb",
                "url": "https://en.wiktionary.org/wiki/Template:lb"
            },
            {
                "name": "Template:tr-noun",
                "url": "https://en.wiktionary.org/wiki/Template:tr-noun"
            },
            {
                "name": "Template:qualifier",
                "url": "https://en.wiktionary.org/wiki/Template:qualifier"
            },
            {
                "name": "Template:der",
                "url": "https://en.wiktionary.org/wiki/Template:der"
            },
            {
                "name": "Template:R:pl:PWN",
                "url": "https://en.wiktionary.org/wiki/Template:R:pl:PWN"
            },
            {
                "name": "Template:pt-IPA",
                "url": "https://en.wiktionary.org/wiki/Template:pt-IPA"
            },
            {
                "name": "Template:wp",
                "url": "https://en.wiktionary.org/wiki/Template:wp"
            },
            {
                "name": "Template:col-auto",
                "url": "https://en.wiktionary.org/wiki/Template:col-auto"
            },
            {
                "name": "Template:a",
                "url": "https://en.wiktionary.org/wiki/Template:a"
            },
            {
                "name": "Template:ellipsis of",
                "url": "https://en.wiktionary.org/wiki/Template:ellipsis_of"
            },
            {
                "name": "Template:top2",
                "url": "https://en.wiktionary.org/wiki/Template:top2"
            },
            {
                "name": "Template:desc",
                "url": "https://en.wiktionary.org/wiki/Template:desc"
            },
            {
                "name": "Template:pt-proper noun",
                "url": "https://en.wiktionary.org/wiki/Template:pt-proper_noun"
            },
            {
                "name": "Template:hypo",
                "url": "https://en.wiktionary.org/wiki/Template:hypo"
            },
            {
                "name": "Template:cite-web",
                "url": "https://en.wiktionary.org/wiki/Template:cite-web"
            },
            {
                "name": "Template:tr-infl-noun-c",
                "url": "https://en.wiktionary.org/wiki/Template:tr-infl-noun-c"
            },
            {
                "name": "Template:trans-bottom",
                "url": "https://en.wiktionary.org/wiki/Template:trans-bottom"
            },
            {
                "name": "Template:audio",
                "url": "https://en.wiktionary.org/wiki/Template:audio"
            },
            {
                "name": "Template:rhymes",
                "url": "https://en.wiktionary.org/wiki/Template:rhymes"
            },
            {
                "name": "Template:t",
                "url": "https://en.wiktionary.org/wiki/Template:t"
            },
            {
                "name": "Template:der-top",
                "url": "https://en.wiktionary.org/wiki/Template:der-top"
            },
            {
                "name": "Template:ca-IPA",
                "url": "https://en.wiktionary.org/wiki/Template:ca-IPA"
            },
            {
                "name": "Template:es-pr",
                "url": "https://en.wiktionary.org/wiki/Template:es-pr"
            },
            {
                "name": "Template:bor",
                "url": "https://en.wiktionary.org/wiki/Template:bor"
            },
            {
                "name": "Template:l",
                "url": "https://en.wiktionary.org/wiki/Template:l"
            },
            {
                "name": "Template:head",
                "url": "https://en.wiktionary.org/wiki/Template:head"
            },
            {
                "name": "Template:place",
                "url": "https://en.wiktionary.org/wiki/Template:place"
            },
            {
                "name": "Template:anagrams",
                "url": "https://en.wiktionary.org/wiki/Template:anagrams"
            },
            {
                "name": "Template:pl-proper noun",
                "url": "https://en.wiktionary.org/wiki/Template:pl-proper_noun"
            },
            {
                "name": "Template:commonscat",
                "url": "https://en.wiktionary.org/wiki/Template:commonscat"
            },
            {
                "name": "Template:ca-proper noun",
                "url": "https://en.wiktionary.org/wiki/Template:ca-proper_noun"
            },
            {
                "name": "Template:compound",
                "url": "https://en.wiktionary.org/wiki/Template:compound"
            },
            {
                "name": "Template:trans-top",
                "url": "https://en.wiktionary.org/wiki/Template:trans-top"
            },
            {
                "name": "Template:sk-IPA",
                "url": "https://en.wiktionary.org/wiki/Template:sk-IPA"
            },
            {
                "name": "Template:m",
                "url": "https://en.wiktionary.org/wiki/Template:m"
            },
            {
                "name": "Template:gloss",
                "url": "https://en.wiktionary.org/wiki/Template:gloss"
            },
            {
                "name": "Template:en-proper noun",
                "url": "https://en.wiktionary.org/wiki/Template:en-proper_noun"
            }
        ],
        "is_part_of": {
            "identifier": "enwiktionary",
            "url": "https://en.wiktionary.org"
        },
        "article_body": {
            "html": "<!DOCTYPE html>\n<html prefix=\"dc: http://purl.org/dc/terms/ mw: http://mediawiki.org/rdf/\" about=\"https://en.wiktionary.org/wiki/Special:Redirect/revision/78276816\"><head prefix=\"mwr: https://en.wiktionary.org/wiki/Special:Redirect/\"><meta property=\"mw:TimeUuid\" content=\"b0a72b60-d6be-11ee-b7e1-499c4badfb44\"/><meta charset=\"utf-8\"/><meta property=\"mw:pageId\" content=\"4344\"/><meta property=\"mw:pageNamespace\" content=\"0\"/><link rel=\"dc:replaces\" resource=\"mwr:revision/77790591\"/><meta property=\"mw:revisionSHA1\" content=\"ebe2e9d557ff347b93ab86b341051dd3c99c5a87\"/><meta property=\"dc:modified\" content=\"2024-02-29T04:54:52.000Z\"/><meta property=\"mw:htmlVersion\" content=\"2.8.0\"/><meta property=\"mw:html:version\" content=\"2.8.0\"/><link rel=\"dc:isVersionOf\" href=\"//en.wiktionary.org/wiki/Montreal\"/><base href=\"//en.wiktionary.org/wiki/\"/><title>Montreal</title><meta property=\"mw:generalModules\" content=\"ext.cite.ux-enhancements|mediawiki.page.media|ext.tmh.player\"/><meta property=\"mw:moduleStyles\" content=\"ext.cite.parsoid.styles|ext.cite.styles|ext.tmh.player.styles\"/><link rel=\"stylesheet\" href=\"/w/load.php?lang=en&amp;modules=ext.cite.parsoid.styles%7Cext.cite.styles%7Cext.tmh.player.styles%7Cmediawiki.skinning.content.parsoid%7Cmediawiki.skinning.interface%7Csite.styles&amp;only=styles&amp;skin=vector\"/><meta http-equiv=\"content-language\" content=\"en\"/><meta http-equiv=\"vary\" content=\"Accept\"/></head><body id=\"mwAA\" lang=\"en\" class=\"mw-content-ltr sitedir-ltr ltr mw-body-content parsoid-body mediawiki mw-parser-output\" dir=\"ltr\"><section data-mw-section-id=\"0\" id=\"mwAQ\"><div class=\"disambig-see-also-2\" about=\"#mwt1\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"also\",\"href\":\"./Template:also\"},\"params\":{\"1\":{\"wt\":\"Montréal\"},\"2\":{\"wt\":\"Mont-real\"}},\"i\":0}}]}' id=\"mwAg\"><i>See also:</i> <b class=\"Latn\"><a rel=\"mw:WikiLink\" href=\"./Montréal\" title=\"Montréal\">Montréal</a></b> <i>and</i> <b class=\"Latn\"><a rel=\"mw:WikiLink\" href=\"./Mont-real\" title=\"Mont-real\">Mont-real</a></b></div>\n<meta property=\"mw:PageProp/toc\" data-mw='{\"autoGenerated\":true}' id=\"mwAw\"/></section><section data-mw-section-id=\"1\" id=\"mwBA\"><h2 id=\"English\">English</h2>\n<div class=\"sister-wikipedia sister-project noprint floatright\" style=\"border: 1px solid #aaa; font-size: 90%; background: #f9f9f9; width: 250px; padding: 4px; text-align: left;\" about=\"#mwt2\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"wikipedia\",\"href\":\"./Template:wikipedia\"},\"params\":{},\"i\":0}}]}' id=\"mwBQ\"><div style=\"float: left;\"><figure class=\"mw-halign-none\" typeof=\"mw:File\"><span><img alt=\"\" resource=\"./File:Wikipedia-logo-v2.svg\" src=\"//upload.wikimedia.org/wikipedia/commons/thumb/8/80/Wikipedia-logo-v2.svg/44px-Wikipedia-logo-v2.svg.png\" decoding=\"async\" data-file-width=\"103\" data-file-height=\"94\" data-file-type=\"drawing\" height=\"40\" width=\"44\" srcset=\"//upload.wikimedia.org/wikipedia/commons/thumb/8/80/Wikipedia-logo-v2.svg/66px-Wikipedia-logo-v2.svg.png 1.5x, //upload.wikimedia.org/wikipedia/commons/thumb/8/80/Wikipedia-logo-v2.svg/88px-Wikipedia-logo-v2.svg.png 2x\" class=\"mw-file-element\"/></span><figcaption></figcaption></figure></div><div style=\"margin-left: 60px;\">English <a rel=\"mw:WikiLink\" href=\"./Wikipedia\" title=\"Wikipedia\">Wikipedia</a> has an article on:<div style=\"margin-left: 10px;\"><b class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink/Interwiki\" href=\"https://en.wikipedia.org/wiki/Montreal\" title=\"w:Montreal\" class=\"extiw\">Montreal</a></b></div></div><span class=\"interProject\"><a rel=\"mw:WikiLink/Interwiki\" href=\"https://en.wikipedia.org/wiki/Montreal\" title=\"w:Montreal\" class=\"extiw\">Wikipedia </a></span></div>\n<div class=\"noprint\" style=\"clear: right; border: solid #aaa 1px; margin: 0 0 1em 1em; font-size: 90%; background: #f9f9f9; width: 250px; padding: 4px; spacing: 0px; text-align: left; float: right;\" about=\"#mwt3\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"commonscat\",\"href\":\"./Template:commonscat\"},\"params\":{},\"i\":0}}]}' id=\"mwBg\">\n<div style=\"float: left;\"><figure class=\"mw-halign-none\" typeof=\"mw:File\"><a href=\"./File:Commons-logo.svg\" class=\"mw-file-description\" title=\"Commons:Category\"><img alt=\"Commons:Category\" resource=\"./File:Commons-logo.svg\" src=\"//upload.wikimedia.org/wikipedia/commons/thumb/4/4a/Commons-logo.svg/50px-Commons-logo.svg.png\" decoding=\"async\" data-file-width=\"1024\" data-file-height=\"1376\" data-file-type=\"drawing\" height=\"67\" width=\"50\" srcset=\"//upload.wikimedia.org/wikipedia/commons/thumb/4/4a/Commons-logo.svg/75px-Commons-logo.svg.png 1.5x, //upload.wikimedia.org/wikipedia/commons/thumb/4/4a/Commons-logo.svg/100px-Commons-logo.svg.png 2x\" class=\"mw-file-element\"/></a><figcaption>Commons:Category</figcaption></figure></div>\n<div style=\"margin-left: 60px;\"><a rel=\"mw:WikiLink/Interwiki\" href=\"https://commons.wikimedia.org/wiki/Main_Page\" title=\"commons:Main Page\" class=\"extiw\">Wikimedia Commons</a> has more media related to:\n<div style=\"margin-left: 10px;\"><i><b><a rel=\"mw:WikiLink/Interwiki\" href=\"https://commons.wikimedia.org/wiki/Category:Montreal\" title=\"commons:Category:Montreal\" class=\"extiw\">Montreal</a></b></i></div>\n</div>\n</div>\n<figure class=\"mw-default-size mw-halign-right\" typeof=\"mw:File/Thumb\" id=\"mwBw\"><a href=\"./File:Montreal_Twilight_Panorama_2006.jpg\" class=\"mw-file-description\" id=\"mwCA\"><img resource=\"./File:Montreal_Twilight_Panorama_2006.jpg\" src=\"//upload.wikimedia.org/wikipedia/commons/thumb/4/47/Montreal_Twilight_Panorama_2006.jpg/220px-Montreal_Twilight_Panorama_2006.jpg\" decoding=\"async\" data-file-width=\"9017\" data-file-height=\"2724\" data-file-type=\"bitmap\" height=\"66\" width=\"220\" srcset=\"//upload.wikimedia.org/wikipedia/commons/thumb/4/47/Montreal_Twilight_Panorama_2006.jpg/330px-Montreal_Twilight_Panorama_2006.jpg 1.5x, //upload.wikimedia.org/wikipedia/commons/thumb/4/47/Montreal_Twilight_Panorama_2006.jpg/440px-Montreal_Twilight_Panorama_2006.jpg 2x\" class=\"mw-file-element\" id=\"mwCQ\"/></a><figcaption id=\"mwCg\">A panorama of Montreal.</figcaption></figure>\n\n<section data-mw-section-id=\"2\" id=\"mwCw\"><h3 id=\"Alternative_forms\">Alternative forms</h3>\n<ul id=\"mwDA\"><li id=\"mwDQ\"><span class=\"ib-brac qualifier-brac\" about=\"#mwt4\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"qualifier\",\"href\":\"./Template:qualifier\"},\"params\":{\"1\":{\"wt\":\"official\"}},\"i\":0}}]}' id=\"mwDg\">(</span><span class=\"ib-content qualifier-content\" about=\"#mwt4\">official</span><span class=\"ib-brac qualifier-brac\" about=\"#mwt4\" id=\"mwDw\">)</span> <span class=\"Latn\" lang=\"en\" about=\"#mwt5\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"alter\",\"href\":\"./Template:alter\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"Montréal\"}},\"i\":0}}]}' id=\"mwEA\"><a rel=\"mw:WikiLink\" href=\"./Montréal#English\" title=\"Montréal\">Montréal</a></span></li></ul>\n\n</section><section data-mw-section-id=\"3\" id=\"mwEQ\"><h3 id=\"Etymology\">Etymology</h3>\n<p id=\"mwEg\">From <span class=\"etyl\" about=\"#mwt7\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"bor\",\"href\":\"./Template:bor\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"fr\"},\"3\":{\"wt\":\"Montréal\"}},\"i\":0}}]}' id=\"mwEw\"><a rel=\"mw:WikiLink/Interwiki\" href=\"https://en.wikipedia.org/wiki/French%20language\" title=\"w:French language\" class=\"extiw\">French</a><link rel=\"mw:PageProp/Category\" href=\"./Category:English_terms_borrowed_from_French#MONTREAL\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:English_terms_derived_from_French#MONTREAL\"/></span><span about=\"#mwt7\"> </span><i class=\"Latn mention\" lang=\"fr\" about=\"#mwt7\" id=\"mwFA\"><a rel=\"mw:WikiLink\" href=\"./Montréal#French\" title=\"Montréal\">Montréal</a></i>; from <i class=\"Latn mention\" lang=\"fr\" about=\"#mwt8\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"m\",\"href\":\"./Template:m\"},\"params\":{\"1\":{\"wt\":\"fr\"},\"2\":{\"wt\":\"Mont Royal\"},\"gloss\":{\"wt\":\"[[Mount Royal]]\"}},\"i\":0}}]}' id=\"mwFQ\"><a rel=\"mw:WikiLink\" href=\"./Mont_Royal?action=edit&amp;redlink=1#French\" title=\"Mont Royal\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Mont Royal\"]}}'>Mont Royal</a></i><span about=\"#mwt8\"> </span><span class=\"mention-gloss-paren annotation-paren\" about=\"#mwt8\">(</span><span class=\"mention-gloss-double-quote\" about=\"#mwt8\">“</span><span class=\"mention-gloss\" about=\"#mwt8\"><a rel=\"mw:WikiLink\" href=\"./Mount_Royal\" title=\"Mount Royal\">Mount Royal</a></span><span class=\"mention-gloss-double-quote\" about=\"#mwt8\">”</span><span class=\"mention-gloss-paren annotation-paren\" about=\"#mwt8\" id=\"mwFg\">)</span> <span class=\"gloss-brac\" about=\"#mwt9\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"gloss\",\"href\":\"./Template:gloss\"},\"params\":{\"1\":{\"wt\":\"{{compound|fr|mont|royal|nocat=1}}\"}},\"i\":0}}]}' id=\"mwFw\">(</span><span class=\"gloss-content\" about=\"#mwt9\"><span class=\"Latn\" lang=\"en\"><i class=\"Latn mention\" lang=\"fr\"><a rel=\"mw:WikiLink\" href=\"./mont#French\" title=\"mont\">mont</a></i> <span typeof=\"mw:Entity\">+</span>‎ <i class=\"Latn mention\" lang=\"fr\"><a rel=\"mw:WikiLink\" href=\"./royal#French\" title=\"royal\">royal</a></i></span></span><span class=\"gloss-brac\" about=\"#mwt9\" id=\"mwGA\">)</span>, after the nearby hill by Jacques Cartier. It is uncertain how <i class=\"Latn mention\" lang=\"fr\" about=\"#mwt10\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"m\",\"href\":\"./Template:m\"},\"params\":{\"1\":{\"wt\":\"fr\"},\"2\":{\"wt\":\"royal\"},\"3\":{\"wt\":\"Royal\"}},\"i\":0}}]}' id=\"mwGQ\"><a rel=\"mw:WikiLink\" href=\"./royal#French\" title=\"royal\">Royal</a></i> became <i id=\"mwGg\">-real</i>. A common explanation is that <i class=\"Latn mention\" lang=\"frm\" about=\"#mwt11\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"m\",\"href\":\"./Template:m\"},\"params\":{\"1\":{\"wt\":\"frm\"},\"2\":{\"wt\":\"real\"}},\"i\":0}}]}' id=\"mwGw\"><a rel=\"mw:WikiLink\" href=\"./real#Middle_French\" title=\"real\">real</a></i> is the <span class=\"etyl\" about=\"#mwt12\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"der\",\"href\":\"./Template:der\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"frm\"},\"3\":{\"wt\":\"-\"}},\"i\":0}}]}' id=\"mwHA\"><a rel=\"mw:WikiLink/Interwiki\" href=\"https://en.wikipedia.org/wiki/Middle%20French\" title=\"w:Middle French\" class=\"extiw\">Middle French</a><link rel=\"mw:PageProp/Category\" href=\"./Category:English_terms_derived_from_Middle_French#MONTREAL\"/></span> form of <i id=\"mwHQ\">Royal</i>, but Cartier himself recorded <i id=\"mwHg\">le mont Royal</i> when he named it. It may have perhaps been from the Italian G.B. Rasmusio's 1556 map translating the name to <span class=\"etyl\" about=\"#mwt13\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"der\",\"href\":\"./Template:der\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"it\"},\"3\":{\"wt\":\"Monte Real\"}},\"i\":0}}]}' id=\"mwHw\"><a rel=\"mw:WikiLink/Interwiki\" href=\"https://en.wikipedia.org/wiki/Italian%20language\" title=\"w:Italian language\" class=\"extiw\">Italian</a><link rel=\"mw:PageProp/Category\" href=\"./Category:English_terms_derived_from_Italian#MONTREAL\"/></span><span about=\"#mwt13\"> </span><i class=\"Latn mention\" lang=\"it\" about=\"#mwt13\" id=\"mwIA\"><a rel=\"mw:WikiLink\" href=\"./Monte_Real?action=edit&amp;redlink=1#Italian\" title=\"Monte Real\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Monte Real\"]}}'>Monte Real</a></i>.<sup about=\"#mwt15\" class=\"mw-ref reference\" id=\"cite_ref-1\" rel=\"dc:references\" typeof=\"mw:Extension/ref\" data-mw='{\"name\":\"ref\",\"attrs\":{},\"body\":{\"id\":\"mw-reference-text-cite_note-1\"}}'><a href=\"./Montreal#cite_note-1\" style=\"counter-reset: mw-Ref 1;\" id=\"mwIQ\"><span class=\"mw-reflink-text\" id=\"mwIg\">[1]</span></a></sup></p>\n\n</section><section data-mw-section-id=\"4\" id=\"mwIw\"><h3 id=\"Pronunciation\">Pronunciation</h3>\n<ul id=\"mwJA\"><li id=\"mwJQ\"><span class=\"ib-brac\" about=\"#mwt16\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"a\",\"href\":\"./Template:a\"},\"params\":{\"1\":{\"wt\":\"UK\"}},\"i\":0}}]}' id=\"mwJg\">(</span><span class=\"ib-content\" about=\"#mwt16\"><a rel=\"mw:WikiLink/Interwiki\" href=\"https://en.wikipedia.org/wiki/British%20English\" title=\"w:British English\" class=\"extiw\">UK</a></span><span class=\"ib-brac\" about=\"#mwt16\" id=\"mwJw\">)</span> <a rel=\"mw:WikiLink\" href=\"./Wiktionary:International_Phonetic_Alphabet\" title=\"Wiktionary:International Phonetic Alphabet\" about=\"#mwt17\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"IPA\",\"href\":\"./Template:IPA\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"/ˌmɒntɹiˈɔːl/\"}},\"i\":0}}]}' id=\"mwKA\">IPA</a><sup about=\"#mwt17\">(<a rel=\"mw:WikiLink\" href=\"./Appendix:English_pronunciation\" title=\"Appendix:English pronunciation\">key</a>)</sup><span about=\"#mwt17\">:</span><span typeof=\"mw:Entity\" about=\"#mwt17\"> </span><span class=\"IPA\" about=\"#mwt17\">/ˌmɒntɹiˈɔːl/</span><link rel=\"mw:PageProp/Category\" href=\"./Category:English_3-syllable_words\" about=\"#mwt17\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:English_terms_with_IPA_pronunciation#MONTREAL\" about=\"#mwt17\" id=\"mwKQ\"/></li>\n<li id=\"mwKg\"><span class=\"ib-brac\" about=\"#mwt18\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"a\",\"href\":\"./Template:a\"},\"params\":{\"1\":{\"wt\":\"US\"}},\"i\":0}}]}' id=\"mwKw\">(</span><span class=\"ib-content\" about=\"#mwt18\"><a rel=\"mw:WikiLink/Interwiki\" href=\"https://en.wikipedia.org/wiki/American%20English\" title=\"w:American English\" class=\"extiw\">US</a></span><span class=\"ib-brac\" about=\"#mwt18\" id=\"mwLA\">)</span> <a rel=\"mw:WikiLink\" href=\"./Wiktionary:International_Phonetic_Alphabet\" title=\"Wiktionary:International Phonetic Alphabet\" about=\"#mwt19\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"IPA\",\"href\":\"./Template:IPA\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"/ˌmɑntɹiˈɔl/\"}},\"i\":0}}]}' id=\"mwLQ\">IPA</a><sup about=\"#mwt19\">(<a rel=\"mw:WikiLink\" href=\"./Appendix:English_pronunciation\" title=\"Appendix:English pronunciation\">key</a>)</sup><span about=\"#mwt19\">:</span><span typeof=\"mw:Entity\" about=\"#mwt19\"> </span><span class=\"IPA\" about=\"#mwt19\">/ˌmɑntɹiˈɔl/</span><link rel=\"mw:PageProp/Category\" href=\"./Category:English_3-syllable_words\" about=\"#mwt19\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:English_terms_with_IPA_pronunciation#MONTREAL\" about=\"#mwt19\" id=\"mwLg\"/></li>\n<li id=\"mwLw\"><span class=\"ib-brac\" about=\"#mwt20\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"a\",\"href\":\"./Template:a\"},\"params\":{\"1\":{\"wt\":\"CA\"}},\"i\":0}}]}' id=\"mwMA\">(</span><span class=\"ib-content\" about=\"#mwt20\"><a rel=\"mw:WikiLink/Interwiki\" href=\"https://en.wikipedia.org/wiki/Canadian%20English\" title=\"w:Canadian English\" class=\"extiw\">Canada</a></span><span class=\"ib-brac\" about=\"#mwt20\" id=\"mwMQ\">)</span> <a rel=\"mw:WikiLink\" href=\"./Wiktionary:International_Phonetic_Alphabet\" title=\"Wiktionary:International Phonetic Alphabet\" about=\"#mwt21\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"IPA\",\"href\":\"./Template:IPA\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"/ˌmʌntɹiːˈɒl/\"},\"3\":{\"wt\":\"/ˌmɔ̃ɹeɪˈæl/\"}},\"i\":0}}]}' id=\"mwMg\">IPA</a><sup about=\"#mwt21\">(<a rel=\"mw:WikiLink\" href=\"./Appendix:English_pronunciation\" title=\"Appendix:English pronunciation\">key</a>)</sup><span about=\"#mwt21\">:</span><span typeof=\"mw:Entity\" about=\"#mwt21\"> </span><span class=\"IPA\" about=\"#mwt21\">/ˌmʌntɹiːˈɒl/</span><span about=\"#mwt21\">, </span><span class=\"IPA\" about=\"#mwt21\">/ˌmɔ̃ɹeɪˈæl/</span><link rel=\"mw:PageProp/Category\" href=\"./Category:English_3-syllable_words\" about=\"#mwt21\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:English_3-syllable_words\" about=\"#mwt21\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:English_terms_with_IPA_pronunciation#MONTREAL\" about=\"#mwt21\" id=\"mwMw\"/>\n<ul id=\"mwNA\"><li id=\"mwNQ\"><style data-mw-deduplicate=\"TemplateStyles:r50165410\" typeof=\"mw:Extension/templatestyles mw:Transclusion\" about=\"#mwt22\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"audio\",\"href\":\"./Template:audio\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"Montreal-English-pronunciation.oga\"},\"3\":{\"wt\":\"Audio (Canada)\"}},\"i\":0}}]}' id=\"mwNg\">.mw-parser-output .k-player .k-attribution{visibility:hidden}</style><table class=\"audiotable\" style=\"vertical-align: middle; display:inline-block; list-style:none;line-height: 1em; border-collapse:collapse;\" about=\"#mwt22\"><tbody><tr><td class=\"unicode audiolink\" style=\"padding-right:5px; padding-left: 0;\">Audio (Canada)</td><td class=\"audiofile\"><span typeof=\"mw:File\" class=\"mw-default-audio-height\" data-mw='{\"caption\":\"noicon\"}'><span title=\"noicon\"><audio controls=\"\" preload=\"none\" height=\"32\" width=\"175\" resource=\"./File:Montreal-English-pronunciation.oga\" data-durationhint=\"1\" class=\"mw-file-element\"><source src=\"//upload.wikimedia.org/wikipedia/commons/5/57/Montreal-English-pronunciation.oga\" type='audio/ogg; codecs=\"vorbis\"'/><source src=\"//upload.wikimedia.org/wikipedia/commons/transcoded/5/57/Montreal-English-pronunciation.oga/Montreal-English-pronunciation.oga.mp3\" type=\"audio/mpeg\" data-transcodekey=\"mp3\"/></audio></span></span></td><td class=\"audiometa\" style=\"font-size: 80%;\">(<a rel=\"mw:WikiLink\" href=\"./File:Montreal-English-pronunciation.oga\" title=\"File:Montreal-English-pronunciation.oga\">file</a>)</td></tr></tbody></table><link rel=\"mw:PageProp/Category\" href=\"./Category:English_terms_with_audio_links#MONTREAL\" about=\"#mwt22\" id=\"mwNw\"/></li></ul></li>\n<li id=\"mwOA\"><span about=\"#mwt25\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"rhymes\",\"href\":\"./Template:rhymes\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"ɔːl\"},\"s\":{\"wt\":\"3\"}},\"i\":0}}]}' id=\"mwOQ\">Rhymes: </span><a rel=\"mw:WikiLink\" href=\"./Rhymes:English/ɔːl\" title=\"Rhymes:English/ɔːl\" about=\"#mwt25\"><span class=\"IPA\">-ɔːl</span></a><link rel=\"mw:PageProp/Category\" href=\"./Category:Rhymes:English/ɔːl#MONTREAL\" about=\"#mwt25\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Rhymes:English/ɔːl/3_syllables#MONTREAL\" about=\"#mwt25\" id=\"mwOg\"/></li></ul>\n\n</section><section data-mw-section-id=\"5\" id=\"mwOw\"><h3 id=\"Proper_noun\">Proper noun</h3>\n<p id=\"mwPA\"><span class=\"headword-line\" about=\"#mwt26\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"en-proper noun\",\"href\":\"./Template:en-proper_noun\"},\"params\":{},\"i\":0}}]}' id=\"mwPQ\"><strong class=\"Latn headword\" lang=\"en\">Montreal</strong></span><link rel=\"mw:PageProp/Category\" href=\"./Category:English_lemmas#MONTREAL\" about=\"#mwt26\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:English_proper_nouns#MONTREAL\" about=\"#mwt26\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:English_uncountable_nouns#MONTREAL\" about=\"#mwt26\" id=\"mwPg\"/></p>\n\n<ol id=\"mwPw\"><li id=\"mwQA\"><span about=\"#mwt27\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"place\",\"href\":\"./Template:place\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"place\"},\"3\":{\"wt\":\"p/Quebec\"},\"4\":{\"wt\":\"c/Canada\"}},\"i\":0}}]}' id=\"mwQQ\">A place in </span><span class=\"Latn\" lang=\"en\" about=\"#mwt27\"><a rel=\"mw:WikiLink\" href=\"./Quebec#English\" title=\"Quebec\">Quebec</a></span><span about=\"#mwt27\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt27\"><a rel=\"mw:WikiLink\" href=\"./Canada#English\" title=\"Canada\">Canada</a></span><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Quebec#MONTREAL\" about=\"#mwt27\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Canada#MONTREAL\" about=\"#mwt27\" id=\"mwQg\"/>, at the confluence of the Ottawa and Saint-Lawrence Rivers\n<ol id=\"mwQw\"><li id=\"mwRA\"><span about=\"#mwt28\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"place\",\"href\":\"./Template:place\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"A river &lt;&lt;port>>, the &lt;&lt;largest city>> in &lt;&lt;p/Quebec>>, &lt;&lt;c/Canada>>\"}},\"i\":0}}]}' id=\"mwRQ\">A river </span><a rel=\"mw:WikiLink\" href=\"./port\" title=\"port\" about=\"#mwt28\">port</a><span about=\"#mwt28\">, the largest city in </span><span class=\"Latn\" lang=\"en\" about=\"#mwt28\"><a rel=\"mw:WikiLink\" href=\"./Quebec#English\" title=\"Quebec\">Quebec</a></span><span about=\"#mwt28\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt28\"><a rel=\"mw:WikiLink\" href=\"./Canada#English\" title=\"Canada\">Canada</a></span><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Montreal#MONTREAL\" about=\"#mwt28\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Cities_in_Quebec#MONTREAL\" about=\"#mwt28\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Cities_in_Canada#MONTREAL\" about=\"#mwt28\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Quebec#MONTREAL\" about=\"#mwt28\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Canada#MONTREAL\" about=\"#mwt28\" id=\"mwRg\"/>.\n<dl id=\"mwRw\"><dd id=\"mwSA\"><span class=\"nyms synonym\" about=\"#mwt29\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"syn\",\"href\":\"./Template:syn\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"City of Montreal\"}},\"i\":0}}]}' id=\"mwSQ\"><span class=\"defdate\">Synonym:</span> <span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./City_of_Montreal?action=edit&amp;redlink=1#English\" title=\"City of Montreal\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"City of Montreal\"]}}'>City of Montreal</a></span></span></dd></dl></li>\n<li id=\"mwSg\"><span about=\"#mwt30\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"place\",\"href\":\"./Template:place\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"island\"},\"3\":{\"wt\":\"p/Quebec\"},\"4\":{\"wt\":\"c/Canada\"}},\"i\":0}}]}' id=\"mwSw\">An island in </span><span class=\"Latn\" lang=\"en\" about=\"#mwt30\"><a rel=\"mw:WikiLink\" href=\"./Quebec#English\" title=\"Quebec\">Quebec</a></span><span about=\"#mwt30\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt30\"><a rel=\"mw:WikiLink\" href=\"./Canada#English\" title=\"Canada\">Canada</a></span><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Islands#MONTREAL\" about=\"#mwt30\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Quebec#MONTREAL\" about=\"#mwt30\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Canada#MONTREAL\" about=\"#mwt30\" id=\"mwTA\"/>, on which the city is situated.\n<dl id=\"mwTQ\"><dd id=\"mwTg\"><span class=\"nyms synonym\" about=\"#mwt31\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"syn\",\"href\":\"./Template:syn\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"Island of Montreal\"},\"3\":{\"wt\":\"Montreal Island\"}},\"i\":0}}]}' id=\"mwTw\"><span class=\"defdate\">Synonyms:</span> <span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./Island_of_Montreal?action=edit&amp;redlink=1#English\" title=\"Island of Montreal\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Island of Montreal\"]}}'>Island of Montreal</a></span>, <span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./Montreal_Island?action=edit&amp;redlink=1#English\" title=\"Montreal Island\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreal Island\"]}}'>Montreal Island</a></span></span></dd>\n<dd id=\"mwUA\"><span class=\"nyms coordinate-term\" about=\"#mwt32\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"cot\",\"href\":\"./Template:cot\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"514\"}},\"i\":0}}]}' id=\"mwUQ\"><span class=\"defdate\">Coordinate term:</span> <span class=\"None\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./514#English\" title=\"514\">514</a></span></span></dd></dl></li>\n<li id=\"mwUg\"><span class=\"form-of-definition use-with-mention\" about=\"#mwt33\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"ellipsis of\",\"href\":\"./Template:ellipsis_of\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"Montreal Archipelago\"}},\"i\":0}}]}' id=\"mwUw\"><a rel=\"mw:WikiLink\" href=\"./Appendix:Glossary#ellipsis\" title=\"Appendix:Glossary\">Ellipsis</a> of <span class=\"form-of-definition-link\"><i class=\"Latn mention\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./Montreal_Archipelago?action=edit&amp;redlink=1#English\" title=\"Montreal Archipelago\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreal Archipelago\"]}}'>Montreal Archipelago</a></i></span></span><span about=\"#mwt33\">.</span><link rel=\"mw:PageProp/Category\" href=\"./Category:English_ellipses#MONTREAL\" about=\"#mwt33\" id=\"mwVA\"/>; <span about=\"#mwt34\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"place\",\"href\":\"./Template:place\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"archipelago\"},\"3\":{\"wt\":\"on the\"},\"4\":{\"wt\":\"river/Saint Lawrence River\"},\"5\":{\"wt\":\"in\"},\"6\":{\"wt\":\"p/Quebec\"},\"7\":{\"wt\":\"c/Canada\"}},\"i\":0}}]}' id=\"mwVQ\">An </span><a rel=\"mw:WikiLink\" href=\"./archipelago\" title=\"archipelago\" about=\"#mwt34\">archipelago</a><span about=\"#mwt34\"> on the </span><span class=\"Latn\" lang=\"en\" about=\"#mwt34\"><a rel=\"mw:WikiLink\" href=\"./Saint_Lawrence_River#English\" title=\"Saint Lawrence River\">Saint Lawrence River</a></span><span about=\"#mwt34\"> in </span><span class=\"Latn\" lang=\"en\" about=\"#mwt34\"><a rel=\"mw:WikiLink\" href=\"./Quebec#English\" title=\"Quebec\">Quebec</a></span><span about=\"#mwt34\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt34\"><a rel=\"mw:WikiLink\" href=\"./Canada#English\" title=\"Canada\">Canada</a></span><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Islands#MONTREAL\" about=\"#mwt34\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Quebec#MONTREAL\" about=\"#mwt34\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Canada#MONTREAL\" about=\"#mwt34\" id=\"mwVg\"/>, which contains the island.\n<dl id=\"mwVw\"><dd id=\"mwWA\"><span class=\"nyms synonym\" about=\"#mwt35\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"syn\",\"href\":\"./Template:syn\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"Hochelaga Archipelago\"}},\"i\":0}}]}' id=\"mwWQ\"><span class=\"defdate\">Synonym:</span> <span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./Hochelaga_Archipelago?action=edit&amp;redlink=1#English\" title=\"Hochelaga Archipelago\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Hochelaga Archipelago\"]}}'>Hochelaga Archipelago</a></span></span></dd></dl></li>\n<li id=\"mwWg\"><span about=\"#mwt36\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"place\",\"href\":\"./Template:place\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"administrative region\"},\"3\":{\"wt\":\"p/Quebec\"},\"4\":{\"wt\":\"c/Canada\"}},\"i\":0}}]}' id=\"mwWw\">An </span><a rel=\"mw:WikiLink\" href=\"./administrative_region\" title=\"administrative region\" about=\"#mwt36\">administrative region</a><span about=\"#mwt36\"> of </span><span class=\"Latn\" lang=\"en\" about=\"#mwt36\"><a rel=\"mw:WikiLink\" href=\"./Quebec#English\" title=\"Quebec\">Quebec</a></span><span about=\"#mwt36\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt36\"><a rel=\"mw:WikiLink\" href=\"./Canada#English\" title=\"Canada\">Canada</a></span><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Quebec#MONTREAL\" about=\"#mwt36\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Canada#MONTREAL\" about=\"#mwt36\" id=\"mwXA\"/>, containing the island.\n<dl id=\"mwXQ\"><dd id=\"mwXg\"><span class=\"nyms synonym\" about=\"#mwt37\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"syn\",\"href\":\"./Template:syn\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"Montreal Region\"},\"3\":{\"wt\":\"Region of Montreal\"},\"4\":{\"wt\":\"Region 6\"},\"5\":{\"wt\":\"Region 06\"}},\"i\":0}}]}' id=\"mwXw\"><span class=\"defdate\">Synonyms:</span> <span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./Montreal_Region?action=edit&amp;redlink=1#English\" title=\"Montreal Region\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreal Region\"]}}'>Montreal Region</a></span>, <span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./Region_of_Montreal?action=edit&amp;redlink=1#English\" title=\"Region of Montreal\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Region of Montreal\"]}}'>Region of Montreal</a></span>, <span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./Region_6?action=edit&amp;redlink=1#English\" title=\"Region 6\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Region 6\"]}}'>Region 6</a></span>, <span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./Region_06?action=edit&amp;redlink=1#English\" title=\"Region 06\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Region 06\"]}}'>Region 06</a></span></span></dd></dl></li>\n<li id=\"mwYA\"><span about=\"#mwt38\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"place\",\"href\":\"./Template:place\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"metropolitan area\"},\"3\":{\"wt\":\"p/Quebec\"},\"4\":{\"wt\":\"c/Canada\"}},\"i\":0}}]}' id=\"mwYQ\">A metropolitan area in </span><span class=\"Latn\" lang=\"en\" about=\"#mwt38\"><a rel=\"mw:WikiLink\" href=\"./Quebec#English\" title=\"Quebec\">Quebec</a></span><span about=\"#mwt38\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt38\"><a rel=\"mw:WikiLink\" href=\"./Canada#English\" title=\"Canada\">Canada</a></span><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Quebec#MONTREAL\" about=\"#mwt38\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Canada#MONTREAL\" about=\"#mwt38\" id=\"mwYg\"/>, containing the archipelago.\n<dl id=\"mwYw\"><dd id=\"mwZA\"><span class=\"nyms hyponym\" about=\"#mwt39\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"hypo\",\"href\":\"./Template:hypo\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"Montreal Metropolitan Community\"},\"3\":{\"wt\":\"Montreal Urban Community\"},\"4\":{\"wt\":\"Agglomeration of Montreal\"},\"5\":{\"wt\":\"Greater Montreal Area\"},\"6\":{\"wt\":\"Greater Montreal\"},\"7\":{\"wt\":\"Metro Montreal\"}},\"i\":0}}]}' id=\"mwZQ\"><span class=\"defdate\">Hyponyms:</span> <span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./Montreal_Metropolitan_Community#English\" title=\"Montreal Metropolitan Community\">Montreal Metropolitan Community</a></span>, <span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./Montreal_Urban_Community?action=edit&amp;redlink=1#English\" title=\"Montreal Urban Community\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreal Urban Community\"]}}'>Montreal Urban Community</a></span>, <span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./Agglomeration_of_Montreal?action=edit&amp;redlink=1#English\" title=\"Agglomeration of Montreal\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Agglomeration of Montreal\"]}}'>Agglomeration of Montreal</a></span>, <span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./Greater_Montreal_Area?action=edit&amp;redlink=1#English\" title=\"Greater Montreal Area\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Greater Montreal Area\"]}}'>Greater Montreal Area</a></span>, <span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./Greater_Montreal?action=edit&amp;redlink=1#English\" title=\"Greater Montreal\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Greater Montreal\"]}}'>Greater Montreal</a></span>, <span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./Metro_Montreal?action=edit&amp;redlink=1#English\" title=\"Metro Montreal\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Metro Montreal\"]}}'>Metro Montreal</a></span></span></dd></dl></li>\n<li id=\"mwZg\"><span class=\"usage-label-sense\" about=\"#mwt40\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"lb\",\"href\":\"./Template:lb\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"historical\"}},\"i\":0}}]}' id=\"mwZw\"><span class=\"ib-brac\">(</span><span class=\"ib-content\"><a rel=\"mw:WikiLink\" href=\"./Appendix:Glossary#historical\" title=\"Appendix:Glossary\">historical</a><link rel=\"mw:PageProp/Category\" href=\"./Category:English_terms_with_historical_senses#MONTREAL\"/></span><span class=\"ib-brac\">)</span></span> <span about=\"#mwt41\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"place\",\"href\":\"./Template:place\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"former county\"},\"3\":{\"wt\":\"p/Quebec\"},\"4\":{\"wt\":\"c/Canada\"}},\"i\":0}}]}' id=\"mwaA\">A former county of </span><span class=\"Latn\" lang=\"en\" about=\"#mwt41\"><a rel=\"mw:WikiLink\" href=\"./Quebec#English\" title=\"Quebec\">Quebec</a></span><span about=\"#mwt41\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt41\"><a rel=\"mw:WikiLink\" href=\"./Canada#English\" title=\"Canada\">Canada</a></span><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Historical_political_subdivisions#MONTREAL\" about=\"#mwt41\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Quebec#MONTREAL\" about=\"#mwt41\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Canada#MONTREAL\" about=\"#mwt41\" id=\"mwaQ\"/>, containing the archipelago.\n<dl id=\"mwag\"><dd id=\"mwaw\"><span class=\"nyms synonym\" about=\"#mwt42\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"syn\",\"href\":\"./Template:syn\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"County of Montreal\"},\"3\":{\"wt\":\"Montreal County\"}},\"i\":0}}]}' id=\"mwbA\"><span class=\"defdate\">Synonyms:</span> <span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./County_of_Montreal?action=edit&amp;redlink=1#English\" title=\"County of Montreal\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"County of Montreal\"]}}'>County of Montreal</a></span>, <span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./Montreal_County?action=edit&amp;redlink=1#English\" title=\"Montreal County\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreal County\"]}}'>Montreal County</a></span></span></dd></dl></li>\n<li id=\"mwbQ\"><span about=\"#mwt43\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"place\",\"href\":\"./Template:place\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"district\"},\"3\":{\"wt\":\"p/Quebec\"},\"4\":{\"wt\":\"c/Canada\"}},\"i\":0}}]}' id=\"mwbg\">A </span><a rel=\"mw:WikiLink\" href=\"./district\" title=\"district\" about=\"#mwt43\">district</a><span about=\"#mwt43\"> of </span><span class=\"Latn\" lang=\"en\" about=\"#mwt43\"><a rel=\"mw:WikiLink\" href=\"./Quebec#English\" title=\"Quebec\">Quebec</a></span><span about=\"#mwt43\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt43\"><a rel=\"mw:WikiLink\" href=\"./Canada#English\" title=\"Canada\">Canada</a></span><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Quebec#MONTREAL\" about=\"#mwt43\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Canada#MONTREAL\" about=\"#mwt43\" id=\"mwbw\"/>, containing the county.\n<dl id=\"mwcA\"><dd id=\"mwcQ\"><span class=\"nyms synonym\" about=\"#mwt44\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"syn\",\"href\":\"./Template:syn\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"District of Montreal\"},\"3\":{\"wt\":\"Montreal District\"}},\"i\":0}}]}' id=\"mwcg\"><span class=\"defdate\">Synonyms:</span> <span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./District_of_Montreal?action=edit&amp;redlink=1#English\" title=\"District of Montreal\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"District of Montreal\"]}}'>District of Montreal</a></span>, <span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./Montreal_District?action=edit&amp;redlink=1#English\" title=\"Montreal District\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreal District\"]}}'>Montreal District</a></span></span></dd></dl>\n<ol id=\"mwcw\"><li id=\"mwdA\"><span class=\"usage-label-sense\" about=\"#mwt45\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"lb\",\"href\":\"./Template:lb\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"judicial\"}},\"i\":0}}]}' id=\"mwdQ\"><span class=\"ib-brac\">(</span><span class=\"ib-content\">judicial</span><span class=\"ib-brac\">)</span></span> <span about=\"#mwt46\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"place\",\"href\":\"./Template:place\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"judicial district\"},\"3\":{\"wt\":\"p/Quebec\"},\"4\":{\"wt\":\"c/Canada\"}},\"i\":0}}]}' id=\"mwdg\">A judicial district in </span><span class=\"Latn\" lang=\"en\" about=\"#mwt46\"><a rel=\"mw:WikiLink\" href=\"./Quebec#English\" title=\"Quebec\">Quebec</a></span><span about=\"#mwt46\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt46\"><a rel=\"mw:WikiLink\" href=\"./Canada#English\" title=\"Canada\">Canada</a></span><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Quebec#MONTREAL\" about=\"#mwt46\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Canada#MONTREAL\" about=\"#mwt46\" id=\"mwdw\"/></li>\n<li id=\"mweA\"><span class=\"usage-label-sense\" about=\"#mwt47\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"lb\",\"href\":\"./Template:lb\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"historical\"}},\"i\":0}}]}' id=\"mweQ\"><span class=\"ib-brac\">(</span><span class=\"ib-content\"><a rel=\"mw:WikiLink\" href=\"./Appendix:Glossary#historical\" title=\"Appendix:Glossary\">historical</a><link rel=\"mw:PageProp/Category\" href=\"./Category:English_terms_with_historical_senses#MONTREAL\"/></span><span class=\"ib-brac\">)</span></span> <span about=\"#mwt48\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"place\",\"href\":\"./Template:place\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"former administrative district\"},\"3\":{\"wt\":\"p/Quebec\"},\"4\":{\"wt\":\"c/Canada\"}},\"i\":0}}]}' id=\"mweg\">A former </span><a rel=\"mw:WikiLink/Interwiki\" href=\"https://en.wikipedia.org/wiki/administrative%20district\" title=\"w:administrative district\" about=\"#mwt48\" class=\"extiw\">administrative district</a><span about=\"#mwt48\"> in </span><span class=\"Latn\" lang=\"en\" about=\"#mwt48\"><a rel=\"mw:WikiLink\" href=\"./Quebec#English\" title=\"Quebec\">Quebec</a></span><span about=\"#mwt48\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt48\"><a rel=\"mw:WikiLink\" href=\"./Canada#English\" title=\"Canada\">Canada</a></span><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Quebec#MONTREAL\" about=\"#mwt48\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Canada#MONTREAL\" about=\"#mwt48\" id=\"mwew\"/></li></ol></li>\n<li id=\"mwfA\"><span class=\"usage-label-sense\" about=\"#mwt49\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"lb\",\"href\":\"./Template:lb\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"historical\"}},\"i\":0}}]}' id=\"mwfQ\"><span class=\"ib-brac\">(</span><span class=\"ib-content\"><a rel=\"mw:WikiLink\" href=\"./Appendix:Glossary#historical\" title=\"Appendix:Glossary\">historical</a><link rel=\"mw:PageProp/Category\" href=\"./Category:English_terms_with_historical_senses#MONTREAL\"/></span><span class=\"ib-brac\">)</span></span> <span about=\"#mwt50\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"place\",\"href\":\"./Template:place\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"district\"},\"3\":{\"wt\":\"p/Quebec\"},\"4\":{\"wt\":\"cc/British North America\"}},\"i\":0}}]}' id=\"mwfg\">A </span><a rel=\"mw:WikiLink\" href=\"./district\" title=\"district\" about=\"#mwt50\">district</a><span about=\"#mwt50\"> of </span><span class=\"Latn\" lang=\"en\" about=\"#mwt50\"><a rel=\"mw:WikiLink\" href=\"./Quebec#English\" title=\"Quebec\">Quebec</a></span><span about=\"#mwt50\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt50\"><a rel=\"mw:WikiLink\" href=\"./British_North_America?action=edit&amp;redlink=1#English\" title=\"British North America\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"British North America\"]}}'>British North America</a></span><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Quebec#MONTREAL\" about=\"#mwt50\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Canada#MONTREAL\" about=\"#mwt50\" id=\"mwfw\"/></li>\n<li id=\"mwgA\"><span class=\"usage-label-sense\" about=\"#mwt51\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"lb\",\"href\":\"./Template:lb\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"historical\"}},\"i\":0}}]}' id=\"mwgQ\"><span class=\"ib-brac\">(</span><span class=\"ib-content\"><a rel=\"mw:WikiLink\" href=\"./Appendix:Glossary#historical\" title=\"Appendix:Glossary\">historical</a><link rel=\"mw:PageProp/Category\" href=\"./Category:English_terms_with_historical_senses#MONTREAL\"/></span><span class=\"ib-brac\">)</span></span> <span about=\"#mwt52\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"place\",\"href\":\"./Template:place\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"district\"},\"3\":{\"wt\":\"p/Canada\"},\"4\":{\"wt\":\"cc/New France\"}},\"i\":0}}]}' id=\"mwgg\">A </span><a rel=\"mw:WikiLink\" href=\"./district\" title=\"district\" about=\"#mwt52\">district</a><span about=\"#mwt52\"> of </span><span class=\"Latn\" lang=\"en\" about=\"#mwt52\"><a rel=\"mw:WikiLink\" href=\"./Canada#English\" title=\"Canada\">Canada</a></span><span about=\"#mwt52\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt52\" id=\"mwgw\"><a rel=\"mw:WikiLink\" href=\"./New_France#English\" title=\"New France\">New France</a></span></li></ol></li>\n<li id=\"mwhA\"><span about=\"#mwt53\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"place\",\"href\":\"./Template:place\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"river\"},\"3\":{\"wt\":\"dist/Algoma,Sudbury\"},\"4\":{\"wt\":\"p/Ontario\"},\"5\":{\"wt\":\"c/Canada\"}},\"i\":0}}]}' id=\"mwhQ\">A river in </span><span class=\"Latn\" lang=\"en\" about=\"#mwt53\"><a rel=\"mw:WikiLink\" href=\"./Algoma#English\" title=\"Algoma\">Algoma</a></span><span about=\"#mwt53\"> and </span><span class=\"Latn\" lang=\"en\" about=\"#mwt53\"><a rel=\"mw:WikiLink\" href=\"./Sudbury#English\" title=\"Sudbury\">Sudbury</a></span><span about=\"#mwt53\"> districts, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt53\"><a rel=\"mw:WikiLink\" href=\"./Ontario#English\" title=\"Ontario\">Ontario</a></span><span about=\"#mwt53\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt53\"><a rel=\"mw:WikiLink\" href=\"./Canada#English\" title=\"Canada\">Canada</a></span><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Rivers_in_Ontario#MONTREAL\" about=\"#mwt53\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Rivers_in_Canada#MONTREAL\" about=\"#mwt53\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Ontario#MONTREAL\" about=\"#mwt53\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Canada#MONTREAL\" about=\"#mwt53\" id=\"mwhg\"/>.</li>\n<li id=\"mwhw\"><span about=\"#mwt54\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"place\",\"href\":\"./Template:place\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"river\"},\"3\":{\"wt\":\"dist/Timiskaming\"},\"4\":{\"wt\":\"p/Ontario\"},\"5\":{\"wt\":\"c/Canada\"}},\"i\":0}}]}' id=\"mwiA\">A river in </span><span class=\"Latn\" lang=\"en\" about=\"#mwt54\"><a rel=\"mw:WikiLink\" href=\"./Timiskaming#English\" title=\"Timiskaming\">Timiskaming</a></span><span about=\"#mwt54\"> district, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt54\"><a rel=\"mw:WikiLink\" href=\"./Ontario#English\" title=\"Ontario\">Ontario</a></span><span about=\"#mwt54\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt54\"><a rel=\"mw:WikiLink\" href=\"./Canada#English\" title=\"Canada\">Canada</a></span><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Rivers_in_Ontario#MONTREAL\" about=\"#mwt54\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Rivers_in_Canada#MONTREAL\" about=\"#mwt54\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Ontario#MONTREAL\" about=\"#mwt54\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Canada#MONTREAL\" about=\"#mwt54\" id=\"mwiQ\"/>.</li>\n<li id=\"mwig\"><span about=\"#mwt55\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"place\",\"href\":\"./Template:place\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"river\"},\"3\":{\"wt\":\"cdiv/No. 18\"},\"4\":{\"wt\":\"p/Saskatchewan\"},\"5\":{\"wt\":\"c/Canada\"}},\"i\":0}}]}' id=\"mwiw\">A river in </span><span class=\"Latn\" lang=\"en\" about=\"#mwt55\"><a rel=\"mw:WikiLink\" href=\"./No._18?action=edit&amp;redlink=1#English\" title=\"No. 18\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"No. 18\"]}}'>No. 18</a></span><span about=\"#mwt55\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt55\"><a rel=\"mw:WikiLink\" href=\"./Saskatchewan#English\" title=\"Saskatchewan\">Saskatchewan</a></span><span about=\"#mwt55\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt55\"><a rel=\"mw:WikiLink\" href=\"./Canada#English\" title=\"Canada\">Canada</a></span><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Rivers_in_Saskatchewan#MONTREAL\" about=\"#mwt55\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Rivers_in_Canada#MONTREAL\" about=\"#mwt55\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Saskatchewan#MONTREAL\" about=\"#mwt55\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Canada#MONTREAL\" about=\"#mwt55\" id=\"mwjA\"/>.</li>\n<li id=\"mwjQ\"><span about=\"#mwt56\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"place\",\"href\":\"./Template:place\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"unincorporated community\"},\"3\":{\"wt\":\"co/Camden County\"},\"4\":{\"wt\":\"s/Missouri\"},\"5\":{\"wt\":\"c/USA\"}},\"i\":0}}]}' id=\"mwjg\">An </span><a rel=\"mw:WikiLink\" href=\"./unincorporated\" title=\"unincorporated\" about=\"#mwt56\">unincorporated</a><span about=\"#mwt56\"> </span><a rel=\"mw:WikiLink\" href=\"./community\" title=\"community\" about=\"#mwt56\">community</a><span about=\"#mwt56\"> in </span><span class=\"Latn\" lang=\"en\" about=\"#mwt56\"><a rel=\"mw:WikiLink\" href=\"./Camden_County#English\" title=\"Camden County\">Camden County</a></span><span about=\"#mwt56\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt56\"><a rel=\"mw:WikiLink\" href=\"./Missouri#English\" title=\"Missouri\">Missouri</a></span><span about=\"#mwt56\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt56\"><a rel=\"mw:WikiLink\" href=\"./United_States#English\" title=\"United States\">United States</a></span><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Unincorporated_communities_in_Missouri,_USA#MONTREAL\" about=\"#mwt56\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Unincorporated_communities_in_the_United_States#MONTREAL\" about=\"#mwt56\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Missouri,_USA#MONTREAL\" about=\"#mwt56\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_the_United_States#MONTREAL\" about=\"#mwt56\" id=\"mwjw\"/>.</li>\n<li id=\"mwkA\"><span about=\"#mwt57\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"place\",\"href\":\"./Template:place\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"small city\"},\"3\":{\"wt\":\"co/Iron County\"},\"4\":{\"wt\":\"s/Wisconsin\"},\"5\":{\"wt\":\"c/USA\"}},\"i\":0}}]}' id=\"mwkQ\">A small city in </span><span class=\"Latn\" lang=\"en\" about=\"#mwt57\"><a rel=\"mw:WikiLink\" href=\"./Iron_County#English\" title=\"Iron County\">Iron County</a></span><span about=\"#mwt57\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt57\"><a rel=\"mw:WikiLink\" href=\"./Wisconsin#English\" title=\"Wisconsin\">Wisconsin</a></span><span about=\"#mwt57\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt57\"><a rel=\"mw:WikiLink\" href=\"./United_States#English\" title=\"United States\">United States</a></span><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Montreal#MONTREAL\" about=\"#mwt57\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Cities_in_Wisconsin,_USA#MONTREAL\" about=\"#mwt57\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Cities_in_the_United_States#MONTREAL\" about=\"#mwt57\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Wisconsin,_USA#MONTREAL\" about=\"#mwt57\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_the_United_States#MONTREAL\" about=\"#mwt57\" id=\"mwkg\"/>.</li>\n<li id=\"mwkw\"><span about=\"#mwt58\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"place\",\"href\":\"./Template:place\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"river\"},\"3\":{\"wt\":\"on the\"},\"4\":{\"wt\":\"pen:Suf/Keweenaw\"},\"5\":{\"wt\":\"s/Michigan\"},\"6\":{\"wt\":\"c/USA\"}},\"i\":0}}]}' id=\"mwlA\">A river on the </span><span class=\"Latn\" lang=\"en\" about=\"#mwt58\"><a rel=\"mw:WikiLink\" href=\"./Keweenaw?action=edit&amp;redlink=1#English\" title=\"Keweenaw\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Keweenaw\"]}}'>Keweenaw</a></span><span about=\"#mwt58\"> Peninsula, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt58\"><a rel=\"mw:WikiLink\" href=\"./Michigan#English\" title=\"Michigan\">Michigan</a></span><span about=\"#mwt58\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt58\"><a rel=\"mw:WikiLink\" href=\"./United_States#English\" title=\"United States\">United States</a></span><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Rivers_in_Michigan,_USA#MONTREAL\" about=\"#mwt58\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Rivers_in_the_United_States#MONTREAL\" about=\"#mwt58\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Michigan,_USA#MONTREAL\" about=\"#mwt58\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_the_United_States#MONTREAL\" about=\"#mwt58\" id=\"mwlQ\"/>.</li>\n<li id=\"mwlg\"><span about=\"#mwt59\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"place\",\"href\":\"./Template:place\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"A &lt;&lt;river>> in northern &lt;&lt;s/Wisconsin>> and the &lt;&lt;pen/Upper Peninsula>> of &lt;&lt;s/Michigan>>, &lt;&lt;c/USA>>\"}},\"i\":0}}]}' id=\"mwlw\">A river in northern </span><span class=\"Latn\" lang=\"en\" about=\"#mwt59\"><a rel=\"mw:WikiLink\" href=\"./Wisconsin#English\" title=\"Wisconsin\">Wisconsin</a></span><span about=\"#mwt59\"> and the </span><span class=\"Latn\" lang=\"en\" about=\"#mwt59\"><a rel=\"mw:WikiLink\" href=\"./Upper_Peninsula#English\" title=\"Upper Peninsula\">Upper Peninsula</a></span><span about=\"#mwt59\"> of </span><span class=\"Latn\" lang=\"en\" about=\"#mwt59\"><a rel=\"mw:WikiLink\" href=\"./Michigan#English\" title=\"Michigan\">Michigan</a></span><span about=\"#mwt59\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt59\"><a rel=\"mw:WikiLink\" href=\"./United_States#English\" title=\"United States\">United States</a></span><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Rivers_in_Wisconsin,_USA#MONTREAL\" about=\"#mwt59\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Rivers_in_the_United_States#MONTREAL\" about=\"#mwt59\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Rivers_in_Michigan,_USA#MONTREAL\" about=\"#mwt59\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Rivers_in_the_United_States#MONTREAL\" about=\"#mwt59\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Wisconsin,_USA#MONTREAL\" about=\"#mwt59\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_the_United_States#MONTREAL\" about=\"#mwt59\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_Michigan,_USA#MONTREAL\" about=\"#mwt59\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:en:Places_in_the_United_States#MONTREAL\" about=\"#mwt59\" id=\"mwmA\"/>.</li></ol>\n\n<section data-mw-section-id=\"6\" id=\"mwmQ\"><h4 id=\"Synonyms\">Synonyms</h4>\n<ul id=\"mwmg\"><li id=\"mwmw\"><span class=\"Latn\" lang=\"en\" about=\"#mwt60\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"l\",\"href\":\"./Template:l\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"Mtl.\"}},\"i\":0}}]}' id=\"mwnA\"><a rel=\"mw:WikiLink\" href=\"./Mtl.#English\" title=\"Mtl.\">Mtl.</a></span>, <span class=\"Latn\" lang=\"en\" about=\"#mwt61\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"l\",\"href\":\"./Template:l\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"Mtl\"}},\"i\":0}}]}' id=\"mwnQ\"><a rel=\"mw:WikiLink\" href=\"./Mtl#English\" title=\"Mtl\">Mtl</a></span>, <span class=\"Latn\" lang=\"en\" about=\"#mwt62\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"l\",\"href\":\"./Template:l\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"MTL\"}},\"i\":0}}]}' id=\"mwng\"><a rel=\"mw:WikiLink\" href=\"./MTL#English\" title=\"MTL\">MTL</a></span>; <span class=\"gloss-brac\" about=\"#mwt63\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"gloss\",\"href\":\"./Template:gloss\"},\"params\":{\"1\":{\"wt\":\"abbreviation\"}},\"i\":0}}]}' id=\"mwnw\">(</span><span class=\"gloss-content\" about=\"#mwt63\"><span class=\"Latn\" lang=\"en\">abbreviation</span></span><span class=\"gloss-brac\" about=\"#mwt63\" id=\"mwoA\">)</span></li></ul>\n\n</section><section data-mw-section-id=\"7\" id=\"mwoQ\"><h4 id=\"Derived_terms\">Derived terms</h4>\n<div class=\"NavFrame\" about=\"#mwt64\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"der-top\",\"href\":\"./Template:der-top\"},\"params\":{},\"i\":0}},\"\\n* \",{\"template\":{\"target\":{\"wt\":\"l\",\"href\":\"./Template:l\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"Montrealer\"}},\"i\":1}},\"\\n* \",{\"template\":{\"target\":{\"wt\":\"l\",\"href\":\"./Template:l\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"Montreal-style\"}},\"i\":2}},\"\\n* \",{\"template\":{\"target\":{\"wt\":\"l\",\"href\":\"./Template:l\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"Montreal bagel\"}},\"i\":3}},\"\\n* \",{\"template\":{\"target\":{\"wt\":\"l\",\"href\":\"./Template:l\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"Montreal smoked meat\"}},\"i\":4}},\"\\n* \",{\"template\":{\"target\":{\"wt\":\"l\",\"href\":\"./Template:l\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"Montreal staircase\"}},\"i\":5}},\"\\n* \",{\"template\":{\"target\":{\"wt\":\"l\",\"href\":\"./Template:l\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"Montreal North\"}},\"i\":6}},\"\\n* \",{\"template\":{\"target\":{\"wt\":\"l\",\"href\":\"./Template:l\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"Montreal nutmeg melon\"}},\"i\":7}},\"\\n*\",{\"template\":{\"target\":{\"wt\":\"l\",\"href\":\"./Template:l\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"Montreal Protocol\"}},\"i\":8}},\"\\n* \",{\"template\":{\"target\":{\"wt\":\"l\",\"href\":\"./Template:l\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"Montreal South\"}},\"i\":9}},\"\\n* \",{\"template\":{\"target\":{\"wt\":\"l\",\"href\":\"./Template:l\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"Montreal East\"}},\"i\":10}},\"\\n* \",{\"template\":{\"target\":{\"wt\":\"l\",\"href\":\"./Template:l\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"Montreal West\"}},\"i\":11}},\"\\n* \",{\"template\":{\"target\":{\"wt\":\"l\",\"href\":\"./Template:l\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"Old Montreal\"}},\"i\":12}},\"\\n\",{\"template\":{\"target\":{\"wt\":\"der-bottom\",\"href\":\"./Template:der-bottom\"},\"params\":{},\"i\":13}}]}' id=\"mwog\">\n<div class=\"NavHead\" style=\"text-align: left;\">Derived terms</div>\n<div class=\"NavContent derivedterms ul-column-count\" data-column-count=\"2\" style=\"text-align: left;\">\n<ul><li><span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./Montrealer#English\" title=\"Montrealer\">Montrealer</a></span></li>\n<li><span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./Montreal-style?action=edit&amp;redlink=1#English\" title=\"Montreal-style\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreal-style\"]}}'>Montreal-style</a></span></li>\n<li><span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./Montreal_bagel?action=edit&amp;redlink=1#English\" title=\"Montreal bagel\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreal bagel\"]}}'>Montreal bagel</a></span></li>\n<li><span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./Montreal_smoked_meat?action=edit&amp;redlink=1#English\" title=\"Montreal smoked meat\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreal smoked meat\"]}}'>Montreal smoked meat</a></span></li>\n<li><span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./Montreal_staircase?action=edit&amp;redlink=1#English\" title=\"Montreal staircase\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreal staircase\"]}}'>Montreal staircase</a></span></li>\n<li><span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./Montreal_North?action=edit&amp;redlink=1#English\" title=\"Montreal North\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreal North\"]}}'>Montreal North</a></span></li>\n<li><span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./Montreal_nutmeg_melon#English\" title=\"Montreal nutmeg melon\">Montreal nutmeg melon</a></span></li>\n<li><span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./Montreal_Protocol#English\" title=\"Montreal Protocol\">Montreal Protocol</a></span></li>\n<li><span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./Montreal_South?action=edit&amp;redlink=1#English\" title=\"Montreal South\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreal South\"]}}'>Montreal South</a></span></li>\n<li><span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./Montreal_East?action=edit&amp;redlink=1#English\" title=\"Montreal East\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreal East\"]}}'>Montreal East</a></span></li>\n<li><span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./Montreal_West?action=edit&amp;redlink=1#English\" title=\"Montreal West\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreal West\"]}}'>Montreal West</a></span></li>\n<li><span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./Old_Montreal?action=edit&amp;redlink=1#English\" title=\"Old Montreal\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Old Montreal\"]}}'>Old Montreal</a></span></li></ul>\n</div></div>\n\n</section><section data-mw-section-id=\"8\" id=\"mwow\"><h4 id=\"Translations\">Translations</h4>\n<div class=\"NavFrame\" id=\"Translations-a_city_in_Quebec\" about=\"#mwt78\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"trans-top\",\"href\":\"./Template:trans-top\"},\"params\":{\"1\":{\"wt\":\"a city in Quebec\"}},\"i\":0}},\"\\n* Arabic: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"ar\"},\"2\":{\"wt\":\"مُونْتْرِيَال\"},\"3\":{\"wt\":\"m\"},\"tr\":{\"wt\":\"montreāl\"}},\"i\":1}},\"\\n* Armenian: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"hy\"},\"2\":{\"wt\":\"Մոնրեալ\"}},\"i\":2}},\"\\n* Belarusian: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"be\"},\"2\":{\"wt\":\"Манрэа́ль\"},\"3\":{\"wt\":\"m\"}},\"i\":3}},\"\\n* Bulgarian: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"bg\"},\"2\":{\"wt\":\"Монреа́л\"},\"3\":{\"wt\":\"m\"}},\"i\":4}},\"\\n* Catalan: \",{\"template\":{\"target\":{\"wt\":\"t+\",\"href\":\"./Template:t+\"},\"params\":{\"1\":{\"wt\":\"ca\"},\"2\":{\"wt\":\"Mont-real\"},\"3\":{\"wt\":\"m\"}},\"i\":5}},\"\\n* Chinese:\\n*: Cantonese: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"yue\"},\"2\":{\"wt\":\"滿地可\"},\"tr\":{\"wt\":\"mun5 dei6 ho2\"}},\"i\":6}},\"\\n*: Mandarin: \",{\"template\":{\"target\":{\"wt\":\"t+\",\"href\":\"./Template:t+\"},\"params\":{\"1\":{\"wt\":\"cmn\"},\"2\":{\"wt\":\"蒙特利爾\"},\"tr\":{\"wt\":\"Méngtèlì&apos;ěr\"}},\"i\":7}},\", \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"cmn\"},\"2\":{\"wt\":\"蒙特婁\"},\"tr\":{\"wt\":\"Méngtèlóu\"}},\"i\":8}},\", \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"cmn\"},\"2\":{\"wt\":\"滿地河\"},\"tr\":{\"wt\":\"Mǎndìhé\"}},\"i\":9}},\", \",{\"template\":{\"target\":{\"wt\":\"t+\",\"href\":\"./Template:t+\"},\"params\":{\"1\":{\"wt\":\"cmn\"},\"2\":{\"wt\":\"蒙城\"},\"tr\":{\"wt\":\"Méngchéng\"}},\"i\":10}},\" \",{\"template\":{\"target\":{\"wt\":\"qualifier\",\"href\":\"./Template:qualifier\"},\"params\":{\"1\":{\"wt\":\"short form\"}},\"i\":11}},\"\\n* Esperanto: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"eo\"},\"2\":{\"wt\":\"Montrealo\"}},\"i\":12}},\"\\n* Farefare: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"gur\"},\"2\":{\"wt\":\"Mõrɛallɛ\"}},\"i\":13}},\"\\n* Finnish: \",{\"template\":{\"target\":{\"wt\":\"t+\",\"href\":\"./Template:t+\"},\"params\":{\"1\":{\"wt\":\"fi\"},\"2\":{\"wt\":\"Montreal\"}},\"i\":14}},\"\\n* French: \",{\"template\":{\"target\":{\"wt\":\"t+\",\"href\":\"./Template:t+\"},\"params\":{\"1\":{\"wt\":\"fr\"},\"2\":{\"wt\":\"Montréal\"},\"3\":{\"wt\":\"m\"}},\"i\":15}},\"\\n*: Middle French: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"frm\"},\"2\":{\"wt\":\"Place Royale\"}},\"i\":16}},\", \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"frm\"},\"2\":{\"wt\":\"Ville Marie\"}},\"i\":17}},\"\\n* Georgian: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"ka\"},\"2\":{\"wt\":\"მონრეალი\"}},\"i\":18}},\"\\n* German: \",{\"template\":{\"target\":{\"wt\":\"t+\",\"href\":\"./Template:t+\"},\"params\":{\"1\":{\"wt\":\"de\"},\"2\":{\"wt\":\"Montreal\"},\"3\":{\"wt\":\"n\"}},\"i\":19}},\", \",{\"template\":{\"target\":{\"wt\":\"t+\",\"href\":\"./Template:t+\"},\"params\":{\"1\":{\"wt\":\"de\"},\"2\":{\"wt\":\"Montréal\"},\"3\":{\"wt\":\"n\"}},\"i\":20}},\"\\n* Greek: \",{\"template\":{\"target\":{\"wt\":\"t+\",\"href\":\"./Template:t+\"},\"params\":{\"1\":{\"wt\":\"el\"},\"2\":{\"wt\":\"Μόντρεαλ\"}},\"i\":21}},\", \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"el\"},\"2\":{\"wt\":\"Μονρεάλη\"}},\"i\":22}},\"\\n* Gujarati: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"gu\"},\"2\":{\"wt\":\"મોન્ટ્રીયલ\"}},\"i\":23}},\"\\n* Haitian Creole: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"ht\"},\"2\":{\"wt\":\"Monreyal\"}},\"i\":24}},\"\\n* Hebrew: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"he\"},\"2\":{\"wt\":\"מונטריאול\"}},\"i\":25}},\"\\n* Hindi: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"hi\"},\"2\":{\"wt\":\"मॉन्ट्रियल\"}},\"i\":26}},\"\\n* Interlingua: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"ia\"},\"2\":{\"wt\":\"Montreal\"}},\"i\":27}},\"\\n* Inuktitut: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"iu\"},\"2\":{\"wt\":\"ᒧᓐᕆᐊᓪ\"},\"sc\":{\"wt\":\"Cans\"}},\"i\":28}},\", \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"iu\"},\"2\":{\"wt\":\"ᒧᓐᑐᕆᐋᓪ\"},\"sc\":{\"wt\":\"Cans\"}},\"i\":29}},\"\\n* Italian: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"it\"},\"2\":{\"wt\":\"Montreale\"},\"3\":{\"wt\":\"f\"}},\"i\":30}},\"\\n* Japanese: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"ja\"},\"2\":{\"wt\":\"モントリオール\"},\"tr\":{\"wt\":\"Montoriōru\"}},\"i\":31}},\", \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"ja\"},\"2\":{\"wt\":\"モンレアル\"},\"tr\":{\"wt\":\"Monrearu\"}},\"i\":32}},\"\\n* Khmer: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"km\"},\"2\":{\"wt\":\"ម៉ុងរ៉េអាល់\"}},\"i\":33}},\"\\n* Korean: \",{\"template\":{\"target\":{\"wt\":\"t+\",\"href\":\"./Template:t+\"},\"params\":{\"1\":{\"wt\":\"ko\"},\"2\":{\"wt\":\"^몬트리올\"}},\"i\":34}},\"\\n* Latin: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"la\"},\"2\":{\"wt\":\"Mons Rēgius\"},\"3\":{\"wt\":\"m\"}},\"i\":35}},\", \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"la\"},\"2\":{\"wt\":\"Mōns Rēgālis\"},\"3\":{\"wt\":\"m\"}},\"i\":36}},\", \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"la\"},\"2\":{\"wt\":\"Marianopolis\"},\"3\":{\"wt\":\"m\"}},\"i\":37}},\"\\n* Latvian: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"lv\"},\"2\":{\"wt\":\"Monreala\"},\"3\":{\"wt\":\"f\"}},\"i\":38}},\"\\n* Laurentian: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"lre\"},\"2\":{\"wt\":\"Tutonaguy\"}},\"i\":39}},\"\\n* Lithuanian: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"lt\"},\"2\":{\"wt\":\"Monrealis\"}},\"i\":40}},\"\\n* Macedonian: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"mk\"},\"2\":{\"wt\":\"Монтреа́л\"},\"3\":{\"wt\":\"m\"}},\"i\":41}},\"\\n* Norman: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"nrf\"},\"2\":{\"wt\":\"Mountriyal\"}},\"i\":42}},\"\\n* Norwegian:\\n*: Bokmål: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"nb\"},\"2\":{\"wt\":\"Montréal\"}},\"i\":43}},\", \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"nb\"},\"2\":{\"wt\":\"Montreal\"}},\"i\":44}},\"\\n*: Nynorsk: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"nn\"},\"2\":{\"wt\":\"Montréal\"}},\"i\":45}},\", \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"nn\"},\"2\":{\"wt\":\"Montreal\"}},\"i\":46}},\"\\n* Ojibwe: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"oj\"},\"2\":{\"wt\":\"Mooniyaang\"}},\"i\":47}},\"\\n* Ottawa: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"otw\"},\"2\":{\"wt\":\"Moon&apos;yaang\"}},\"i\":48}},\"\\n* Persian: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"fa\"},\"2\":{\"wt\":\"مونترال\"}},\"i\":49}},\"\\n* Polish: \",{\"template\":{\"target\":{\"wt\":\"t+\",\"href\":\"./Template:t+\"},\"params\":{\"1\":{\"wt\":\"pl\"},\"2\":{\"wt\":\"Montreal\"},\"3\":{\"wt\":\"m-in\"}},\"i\":50}},\"\\n* Portuguese: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"pt\"},\"2\":{\"wt\":\"Montreal\"},\"3\":{\"wt\":\"f\"}},\"i\":51}},\", \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"pt\"},\"2\":{\"wt\":\"Montréal\"},\"3\":{\"wt\":\"f\"}},\"i\":52}},\"\\n* Romanian: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"ro\"},\"2\":{\"wt\":\"Montreal\"},\"3\":{\"wt\":\"n\"}},\"i\":53}},\"\\n* Russian: \",{\"template\":{\"target\":{\"wt\":\"t+\",\"href\":\"./Template:t+\"},\"params\":{\"1\":{\"wt\":\"ru\"},\"2\":{\"wt\":\"Монреа́ль\"},\"3\":{\"wt\":\"m\"}},\"i\":54}},\"\\n* Scottish Gaelic: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"gd\"},\"2\":{\"wt\":\"Beinn Rìoghail\"},\"3\":{\"wt\":\"f\"}},\"i\":55}},\"\\n* Serbo-Croatian:\\n*: Cyrillic: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"sh\"},\"2\":{\"wt\":\"Монтреал\"},\"3\":{\"wt\":\"m\"}},\"i\":56}},\"\\n*: Roman: \",{\"template\":{\"target\":{\"wt\":\"t+\",\"href\":\"./Template:t+\"},\"params\":{\"1\":{\"wt\":\"sh\"},\"2\":{\"wt\":\"Montreal\"},\"3\":{\"wt\":\"m\"}},\"i\":57}},\"\\n* Spanish: \",{\"template\":{\"target\":{\"wt\":\"t+\",\"href\":\"./Template:t+\"},\"params\":{\"1\":{\"wt\":\"es\"},\"2\":{\"wt\":\"Montreal\"},\"3\":{\"wt\":\"m\"}},\"i\":58}},\"\\n* Thai: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"th\"},\"2\":{\"wt\":\"มอนทรีออล\"}},\"i\":59}},\"\\n* Turkish: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"tr\"},\"2\":{\"wt\":\"Montreal\"}},\"i\":60}},\"\\n* Ukrainian: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"uk\"},\"2\":{\"wt\":\"Монреа́ль\"},\"3\":{\"wt\":\"m\"}},\"i\":61}},\"\\n* Vietnamese: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"vi\"},\"2\":{\"wt\":\"Mông-rê-an\"}},\"i\":62}},\", \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"vi\"},\"2\":{\"wt\":\"Montréal\"}},\"i\":63}},\"\\n* Yiddish: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"yi\"},\"2\":{\"wt\":\"מאָנטרעאָל\"}},\"i\":64}},\", \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"yi\"},\"2\":{\"wt\":\"מאָנטרעאַל\"}},\"i\":65}},\"\\n\",{\"template\":{\"target\":{\"wt\":\"trans-bottom\",\"href\":\"./Template:trans-bottom\"},\"params\":{},\"i\":66}}]}'><div class=\"NavHead\" style=\"text-align:left;cursor:pointer\">a city in Quebec</div><div class=\"NavContent\"><table class=\"translations\" role=\"presentation\" data-gloss=\"a city in Quebec\" style=\"width:100%\"><tbody><tr><td class=\"translations-cell multicolumn-list\" colspan=\"3\" style=\"background-color:#ffffe0;vertical-align:top;text-align:left\">\n<ul><li>Arabic: <span typeof=\"mw:Entity\">⁧</span><span class=\"Arab\" lang=\"ar\"><a rel=\"mw:WikiLink\" href=\"./مونتريال#Arabic\" title=\"مونتريال\">مُونْتْرِيَال</a></span><span typeof=\"mw:Entity\">⁩</span><span typeof=\"mw:Entity\"> </span><span class=\"gender\"><abbr title=\"masculine gender\">m</abbr></span> <span class=\"mention-gloss-paren annotation-paren\">(</span><span lang=\"ar-Latn\" class=\"tr Latn\">montreāl</span><span class=\"mention-gloss-paren annotation-paren\">)</span><link rel=\"mw:PageProp/Category\" href=\"./Category:Arabic_terms_with_non-redundant_manual_transliterations#MONTREAL\"/></li>\n<li>Armenian: <span class=\"Armn\" lang=\"hy\"><a rel=\"mw:WikiLink\" href=\"./Մոնրեալ?action=edit&amp;redlink=1#Armenian\" title=\"Մոնրեալ\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Մոնրեալ\"]}}'>Մոնրեալ</a></span> <span class=\"mention-gloss-paren annotation-paren\">(</span><span lang=\"hy-Latn\" class=\"tr Latn\">Monreal</span><span class=\"mention-gloss-paren annotation-paren\">)</span></li>\n<li>Belarusian: <span class=\"Cyrl\" lang=\"be\"><a rel=\"mw:WikiLink\" href=\"./Манрэаль?action=edit&amp;redlink=1#Belarusian\" title=\"Манрэаль\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Манрэаль\"]}}'>Манрэа́ль</a></span><span typeof=\"mw:Entity\"> </span><span class=\"gender\"><abbr title=\"masculine gender\">m</abbr></span> <span class=\"mention-gloss-paren annotation-paren\">(</span><span lang=\"be-Latn\" class=\"tr Latn\">Manreálʹ</span><span class=\"mention-gloss-paren annotation-paren\">)</span></li>\n<li>Bulgarian: <span class=\"Cyrl\" lang=\"bg\"><a rel=\"mw:WikiLink\" href=\"./Монреал?action=edit&amp;redlink=1#Bulgarian\" title=\"Монреал\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Монреал\"]}}'>Монреа́л</a></span><span typeof=\"mw:Entity\"> </span><span class=\"gender\"><abbr title=\"masculine gender\">m</abbr></span> <span class=\"mention-gloss-paren annotation-paren\">(</span><span lang=\"bg-Latn\" class=\"tr Latn\">Monreál</span><span class=\"mention-gloss-paren annotation-paren\">)</span></li>\n<li>Catalan: <span class=\"Latn\" lang=\"ca\"><a rel=\"mw:WikiLink\" href=\"./Mont-real#Catalan\" title=\"Mont-real\">Mont-real</a></span><span class=\"tpos\"><span typeof=\"mw:Entity\"> </span><a rel=\"mw:WikiLink/Interwiki\" href=\"https://ca.wiktionary.org/wiki/Mont-real\" title=\"ca:Mont-real\" class=\"extiw\">(ca)</a></span><span typeof=\"mw:Entity\"> </span><span class=\"gender\"><abbr title=\"masculine gender\">m</abbr></span></li>\n<li>Chinese:\n<dl><dd>Cantonese: <span class=\"Hant\" lang=\"yue\"><a rel=\"mw:WikiLink\" href=\"./滿地可#Cantonese\" title=\"滿地可\">滿地可</a></span><span class=\"Zsym mention\" style=\"font-size:100%;\">／</span><span class=\"Hans\" lang=\"yue\"><a rel=\"mw:WikiLink\" href=\"./满地可#Cantonese\" title=\"满地可\">满地可</a></span> <span class=\"mention-gloss-paren annotation-paren\">(</span><span lang=\"yue-Latn\" class=\"tr Latn\">mun<sup>5</sup> dei<sup>6</sup> ho<sup>2</sup></span><span class=\"mention-gloss-paren annotation-paren\">)</span><link rel=\"mw:PageProp/Category\" href=\"./Category:Cantonese_terms_with_redundant_transliterations#MONTREAL\"/></dd>\n<dd>Mandarin: <span class=\"Hant\" lang=\"cmn\"><a rel=\"mw:WikiLink\" href=\"./蒙特利爾#Mandarin\" title=\"蒙特利爾\">蒙特利爾</a></span><span class=\"Zsym mention\" style=\"font-size:100%;\">／</span><span class=\"Hans\" lang=\"cmn\"><a rel=\"mw:WikiLink\" href=\"./蒙特利尔#Mandarin\" title=\"蒙特利尔\">蒙特利尔</a></span><span class=\"tpos\"><span typeof=\"mw:Entity\"> </span><a rel=\"mw:WikiLink/Interwiki\" href=\"https://zh.wiktionary.org/wiki/蒙特利爾\" title=\"zh:蒙特利爾\" class=\"extiw\">(zh)</a></span> <span class=\"mention-gloss-paren annotation-paren\">(</span><span lang=\"cmn-Latn\" class=\"tr Latn\">Méngtèlì'ěr</span><span class=\"mention-gloss-paren annotation-paren\">)</span><link rel=\"mw:PageProp/Category\" href=\"./Category:Mandarin_terms_with_redundant_transliterations#MONTREAL\"/>, <span class=\"Hant\" lang=\"cmn\"><a rel=\"mw:WikiLink\" href=\"./蒙特婁#Mandarin\" title=\"蒙特婁\">蒙特婁</a></span><span class=\"Zsym mention\" style=\"font-size:100%;\">／</span><span class=\"Hans\" lang=\"cmn\"><a rel=\"mw:WikiLink\" href=\"./蒙特娄#Mandarin\" title=\"蒙特娄\">蒙特娄</a></span> <span class=\"mention-gloss-paren annotation-paren\">(</span><span lang=\"cmn-Latn\" class=\"tr Latn\">Méngtèlóu</span><span class=\"mention-gloss-paren annotation-paren\">)</span><link rel=\"mw:PageProp/Category\" href=\"./Category:Mandarin_terms_with_redundant_transliterations#MONTREAL\"/>, <span class=\"Hant\" lang=\"cmn\"><a rel=\"mw:WikiLink\" href=\"./滿地河?action=edit&amp;redlink=1#Mandarin\" title=\"滿地河\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"滿地河\"]}}'>滿地河</a></span><span class=\"Zsym mention\" style=\"font-size:100%;\">／</span><span class=\"Hans\" lang=\"cmn\"><a rel=\"mw:WikiLink\" href=\"./满地河?action=edit&amp;redlink=1#Mandarin\" title=\"满地河\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"满地河\"]}}'>满地河</a></span> <span class=\"mention-gloss-paren annotation-paren\">(</span><span lang=\"cmn-Latn\" class=\"tr Latn\">Mǎndìhé</span><span class=\"mention-gloss-paren annotation-paren\">)</span>, <span class=\"Hani\" lang=\"cmn\"><a rel=\"mw:WikiLink\" href=\"./蒙城#Mandarin\" title=\"蒙城\">蒙城</a></span><span class=\"tpos\"><span typeof=\"mw:Entity\"> </span><a rel=\"mw:WikiLink/Interwiki\" href=\"https://zh.wiktionary.org/wiki/蒙城\" title=\"zh:蒙城\" class=\"extiw\">(zh)</a></span> <span class=\"mention-gloss-paren annotation-paren\">(</span><span lang=\"cmn-Latn\" class=\"tr Latn\">Méngchéng</span><span class=\"mention-gloss-paren annotation-paren\">)</span><link rel=\"mw:PageProp/Category\" href=\"./Category:Mandarin_terms_with_redundant_transliterations#MONTREAL\"/> <span class=\"ib-brac qualifier-brac\">(</span><span class=\"ib-content qualifier-content\">short form</span><span class=\"ib-brac qualifier-brac\">)</span></dd></dl></li>\n<li>Esperanto: <span class=\"Latn\" lang=\"eo\"><a rel=\"mw:WikiLink\" href=\"./Montrealo#Esperanto\" title=\"Montrealo\">Montrealo</a></span></li>\n<li>Farefare: <span class=\"Latn\" lang=\"gur\"><a rel=\"mw:WikiLink\" href=\"./Mõrɛallɛ?action=edit&amp;redlink=1#Farefare\" title=\"Mõrɛallɛ\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Mõrɛallɛ\"]}}'>Mõrɛallɛ</a></span></li>\n<li>Finnish: <span class=\"Latn\" lang=\"fi\"><a rel=\"mw:WikiLink\" href=\"./Montreal#Finnish\" class=\"mw-selflink-fragment\">Montreal</a></span><span class=\"tpos\"><span typeof=\"mw:Entity\"> </span><a rel=\"mw:WikiLink/Interwiki\" href=\"https://fi.wiktionary.org/wiki/Montreal\" title=\"fi:Montreal\" class=\"extiw\">(fi)</a></span></li>\n<li>French: <span class=\"Latn\" lang=\"fr\"><a rel=\"mw:WikiLink\" href=\"./Montréal#French\" title=\"Montréal\">Montréal</a></span><span class=\"tpos\"><span typeof=\"mw:Entity\"> </span><a rel=\"mw:WikiLink/Interwiki\" href=\"https://fr.wiktionary.org/wiki/Montréal\" title=\"fr:Montréal\" class=\"extiw\">(fr)</a></span><span typeof=\"mw:Entity\"> </span><span class=\"gender\"><abbr title=\"masculine gender\">m</abbr></span>\n<dl><dd>Middle French: <span class=\"Latn\" lang=\"frm\"><a rel=\"mw:WikiLink\" href=\"./Place_Royale?action=edit&amp;redlink=1#Middle_French\" title=\"Place Royale\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Place Royale\"]}}'>Place Royale</a></span>, <span class=\"Latn\" lang=\"frm\"><a rel=\"mw:WikiLink\" href=\"./Ville_Marie?action=edit&amp;redlink=1#Middle_French\" title=\"Ville Marie\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Ville Marie\"]}}'>Ville Marie</a></span></dd></dl></li>\n<li>Georgian: <span class=\"Geor\" lang=\"ka\"><a rel=\"mw:WikiLink\" href=\"./მონრეალი#Georgian\" title=\"მონრეალი\">მონრეალი</a></span> <span class=\"mention-gloss-paren annotation-paren\">(</span><span lang=\"ka-Latn\" class=\"tr Latn\">monreali</span><span class=\"mention-gloss-paren annotation-paren\">)</span></li>\n<li>German: <span class=\"Latn\" lang=\"de\"><a rel=\"mw:WikiLink\" href=\"./Montreal#German\" class=\"mw-selflink-fragment\">Montreal</a></span><span class=\"tpos\"><span typeof=\"mw:Entity\"> </span><a rel=\"mw:WikiLink/Interwiki\" href=\"https://de.wiktionary.org/wiki/Montreal\" title=\"de:Montreal\" class=\"extiw\">(de)</a></span><span typeof=\"mw:Entity\"> </span><span class=\"gender\"><abbr title=\"neuter gender\">n</abbr></span>, <span class=\"Latn\" lang=\"de\"><a rel=\"mw:WikiLink\" href=\"./Montréal#German\" title=\"Montréal\">Montréal</a></span><span class=\"tpos\"><span typeof=\"mw:Entity\"> </span><a rel=\"mw:WikiLink/Interwiki\" href=\"https://de.wiktionary.org/wiki/Montréal\" title=\"de:Montréal\" class=\"extiw\">(de)</a></span><span typeof=\"mw:Entity\"> </span><span class=\"gender\"><abbr title=\"neuter gender\">n</abbr></span></li>\n<li>Greek: <span class=\"Grek\" lang=\"el\"><a rel=\"mw:WikiLink\" href=\"./Μόντρεαλ?action=edit&amp;redlink=1#Greek\" title=\"Μόντρεαλ\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Μόντρεαλ\"]}}'>Μόντρεαλ</a></span><span class=\"tpos\"><span typeof=\"mw:Entity\"> </span><a rel=\"mw:WikiLink/Interwiki\" href=\"https://el.wiktionary.org/wiki/Μόντρεαλ\" title=\"el:Μόντρεαλ\" class=\"extiw\">(el)</a></span> <span class=\"mention-gloss-paren annotation-paren\">(</span><span lang=\"el-Latn\" class=\"tr Latn\">Móntreal</span><span class=\"mention-gloss-paren annotation-paren\">)</span>, <span class=\"Grek\" lang=\"el\"><a rel=\"mw:WikiLink\" href=\"./Μονρεάλη?action=edit&amp;redlink=1#Greek\" title=\"Μονρεάλη\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Μονρεάλη\"]}}'>Μονρεάλη</a></span> <span class=\"mention-gloss-paren annotation-paren\">(</span><span lang=\"el-Latn\" class=\"tr Latn\">Monreáli</span><span class=\"mention-gloss-paren annotation-paren\">)</span></li>\n<li>Gujarati: <span class=\"Gujr\" lang=\"gu\"><a rel=\"mw:WikiLink\" href=\"./મોન્ટ્રીયલ?action=edit&amp;redlink=1#Gujarati\" title=\"મોન્ટ્રીયલ\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"મોન્ટ્રીયલ\"]}}'>મોન્ટ્રીયલ</a></span> <span class=\"mention-gloss-paren annotation-paren\">(</span><span lang=\"gu-Latn\" class=\"tr Latn\">monṭrīyal</span><span class=\"mention-gloss-paren annotation-paren\">)</span></li>\n<li>Haitian Creole: <span class=\"Latn\" lang=\"ht\"><a rel=\"mw:WikiLink\" href=\"./Monreyal#Haitian_Creole\" title=\"Monreyal\">Monreyal</a></span></li>\n<li>Hebrew: <span typeof=\"mw:Entity\">⁧</span><span class=\"Hebr\" lang=\"he\"><a rel=\"mw:WikiLink\" href=\"./מונטריאול?action=edit&amp;redlink=1#Hebrew\" title=\"מונטריאול\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"מונטריאול\"]}}'>מונטריאול</a></span><span typeof=\"mw:Entity\">⁩</span></li>\n<li>Hindi: <span class=\"Deva\" lang=\"hi\"><a rel=\"mw:WikiLink\" href=\"./मॉन्ट्रियल?action=edit&amp;redlink=1#Hindi\" title=\"मॉन्ट्रियल\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"मॉन्ट्रियल\"]}}'>मॉन्ट्रियल</a></span> <span class=\"mention-gloss-paren annotation-paren\">(</span><span lang=\"hi-Latn\" class=\"tr Latn\">mŏnṭriyal</span><span class=\"mention-gloss-paren annotation-paren\">)</span></li>\n<li>Interlingua: <span class=\"Latn\" lang=\"ia\"><a rel=\"mw:WikiLink\" href=\"./Montreal#Interlingua\" class=\"mw-selflink-fragment\">Montreal</a></span></li>\n<li>Inuktitut: <span class=\"Cans\" lang=\"iu\"><a rel=\"mw:WikiLink\" href=\"./ᒧᓐᕆᐊᓪ?action=edit&amp;redlink=1#Inuktitut\" title=\"ᒧᓐᕆᐊᓪ\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"ᒧᓐᕆᐊᓪ\"]}}'>ᒧᓐᕆᐊᓪ</a></span> <span class=\"mention-gloss-paren annotation-paren\">(</span><span lang=\"iu-Latn\" class=\"tr Latn\">monrial</span><span class=\"mention-gloss-paren annotation-paren\">)</span><link rel=\"mw:PageProp/Category\" href=\"./Category:Inuktitut_terms_with_redundant_script_codes#MONTREAL\"/>, <span class=\"Cans\" lang=\"iu\"><a rel=\"mw:WikiLink\" href=\"./ᒧᓐᑐᕆᐋᓪ?action=edit&amp;redlink=1#Inuktitut\" title=\"ᒧᓐᑐᕆᐋᓪ\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"ᒧᓐᑐᕆᐋᓪ\"]}}'>ᒧᓐᑐᕆᐋᓪ</a></span> <span class=\"mention-gloss-paren annotation-paren\">(</span><span lang=\"iu-Latn\" class=\"tr Latn\">montoriaal</span><span class=\"mention-gloss-paren annotation-paren\">)</span><link rel=\"mw:PageProp/Category\" href=\"./Category:Inuktitut_terms_with_redundant_script_codes#MONTREAL\"/></li>\n<li>Italian: <span class=\"Latn\" lang=\"it\"><a rel=\"mw:WikiLink\" href=\"./Montreale?action=edit&amp;redlink=1#Italian\" title=\"Montreale\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreale\"]}}'>Montreale</a></span><span typeof=\"mw:Entity\"> </span><span class=\"gender\"><abbr title=\"feminine gender\">f</abbr></span></li>\n<li>Japanese: <span class=\"Jpan\" lang=\"ja\"><a rel=\"mw:WikiLink\" href=\"./モントリオール#Japanese\" title=\"モントリオール\">モントリオール</a></span> <span class=\"mention-gloss-paren annotation-paren\">(</span><span class=\"tr\">Montoriōru</span><span class=\"mention-gloss-paren annotation-paren\">)</span>, <span class=\"Jpan\" lang=\"ja\"><a rel=\"mw:WikiLink\" href=\"./モンレアル#Japanese\" title=\"モンレアル\">モンレアル</a></span> <span class=\"mention-gloss-paren annotation-paren\">(</span><span class=\"tr\">Monrearu</span><span class=\"mention-gloss-paren annotation-paren\">)</span></li>\n<li>Khmer: <span class=\"Khmr\" lang=\"km\"><a rel=\"mw:WikiLink\" href=\"./ម៉ុងរ៉េអាល់#Khmer\" title=\"ម៉ុងរ៉េអាល់\">ម៉ុងរ៉េអាល់</a></span> <span class=\"mention-gloss-paren annotation-paren\">(</span><span lang=\"km-Latn\" class=\"tr Latn\">mongreeʼal</span><span class=\"mention-gloss-paren annotation-paren\">)</span></li>\n<li>Korean: <span class=\"Kore\" lang=\"ko\"><a rel=\"mw:WikiLink\" href=\"./몬트리올#Korean\" title=\"몬트리올\">몬트리올</a></span><span class=\"tpos\"><span typeof=\"mw:Entity\"> </span><a rel=\"mw:WikiLink/Interwiki\" href=\"https://ko.wiktionary.org/wiki/몬트리올\" title=\"ko:몬트리올\" class=\"extiw\">(ko)</a></span> <span class=\"mention-gloss-paren annotation-paren\">(</span><span lang=\"ko-Latn\" class=\"tr Latn\">Monteuriol</span><span class=\"mention-gloss-paren annotation-paren\">)</span></li>\n<li>Latin: <span class=\"Latn\" lang=\"la\"><a rel=\"mw:WikiLink\" href=\"./Mons_Regius?action=edit&amp;redlink=1#Latin\" title=\"Mons Regius\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Mons Regius\"]}}'>Mons Rēgius</a></span><span typeof=\"mw:Entity\"> </span><span class=\"gender\"><abbr title=\"masculine gender\">m</abbr></span>, <span class=\"Latn\" lang=\"la\"><a rel=\"mw:WikiLink\" href=\"./Mons_Regalis?action=edit&amp;redlink=1#Latin\" title=\"Mons Regalis\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Mons Regalis\"]}}'>Mōns Rēgālis</a></span><span typeof=\"mw:Entity\"> </span><span class=\"gender\"><abbr title=\"masculine gender\">m</abbr></span>, <span class=\"Latn\" lang=\"la\"><a rel=\"mw:WikiLink\" href=\"./Marianopolis?action=edit&amp;redlink=1#Latin\" title=\"Marianopolis\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Marianopolis\"]}}'>Marianopolis</a></span><span typeof=\"mw:Entity\"> </span><span class=\"gender\"><abbr title=\"masculine gender\">m</abbr></span></li>\n<li>Latvian: <span class=\"Latn\" lang=\"lv\"><a rel=\"mw:WikiLink\" href=\"./Monreala?action=edit&amp;redlink=1#Latvian\" title=\"Monreala\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Monreala\"]}}'>Monreala</a></span><span typeof=\"mw:Entity\"> </span><span class=\"gender\"><abbr title=\"feminine gender\">f</abbr></span></li>\n<li>Laurentian: <span class=\"Latn\" lang=\"lre\"><a rel=\"mw:WikiLink\" href=\"./Tutonaguy?action=edit&amp;redlink=1#Laurentian\" title=\"Tutonaguy\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Tutonaguy\"]}}'>Tutonaguy</a></span></li>\n<li>Lithuanian: <span class=\"Latn\" lang=\"lt\"><a rel=\"mw:WikiLink\" href=\"./Monrealis?action=edit&amp;redlink=1#Lithuanian\" title=\"Monrealis\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Monrealis\"]}}'>Monrealis</a></span></li>\n<li>Macedonian: <span class=\"Cyrl\" lang=\"mk\"><a rel=\"mw:WikiLink\" href=\"./Монтреал?action=edit&amp;redlink=1#Macedonian\" title=\"Монтреал\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Монтреал\"]}}'>Монтреа́л</a></span><span typeof=\"mw:Entity\"> </span><span class=\"gender\"><abbr title=\"masculine gender\">m</abbr></span> <span class=\"mention-gloss-paren annotation-paren\">(</span><span lang=\"mk-Latn\" class=\"tr Latn\">Montreál</span><span class=\"mention-gloss-paren annotation-paren\">)</span></li>\n<li>Norman: <span class=\"Latn\" lang=\"nrf\"><a rel=\"mw:WikiLink\" href=\"./Mountriyal?action=edit&amp;redlink=1#Norman\" title=\"Mountriyal\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Mountriyal\"]}}'>Mountriyal</a></span></li>\n<li>Norwegian:\n<dl><dd>Bokmål: <span class=\"Latn\" lang=\"nb\"><a rel=\"mw:WikiLink\" href=\"./Montréal#Norwegian_Bokmål\" title=\"Montréal\">Montréal</a></span>, <span class=\"Latn\" lang=\"nb\"><a rel=\"mw:WikiLink\" href=\"./Montreal#Norwegian_Bokmål\" class=\"mw-selflink-fragment\">Montreal</a></span></dd>\n<dd>Nynorsk: <span class=\"Latn\" lang=\"nn\"><a rel=\"mw:WikiLink\" href=\"./Montreal#Norwegian_Nynorsk\" class=\"mw-selflink-fragment\">Montréal</a></span>, <span class=\"Latn\" lang=\"nn\"><a rel=\"mw:WikiLink\" href=\"./Montreal#Norwegian_Nynorsk\" class=\"mw-selflink-fragment\">Montreal</a></span></dd></dl></li>\n<li>Ojibwe: <span class=\"Latn\" lang=\"oj\"><a rel=\"mw:WikiLink\" href=\"./Mooniyaang?action=edit&amp;redlink=1#Ojibwe\" title=\"Mooniyaang\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Mooniyaang\"]}}'>Mooniyaang</a></span></li>\n<li>Ottawa: <span class=\"Latn\" lang=\"otw\"><a rel=\"mw:WikiLink\" href=\"./Moon'yaang?action=edit&amp;redlink=1#Ottawa\" title=\"Moon'yaang\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Moon&apos;yaang\"]}}'>Moon'yaang</a></span></li>\n<li>Persian: <span typeof=\"mw:Entity\">⁧</span><span class=\"Arab fa-Arab\" lang=\"fa\"><a rel=\"mw:WikiLink\" href=\"./مونترال?action=edit&amp;redlink=1#Persian\" title=\"مونترال\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"مونترال\"]}}'>مونترال</a></span><span typeof=\"mw:Entity\">⁩</span></li>\n<li>Polish: <span class=\"Latn\" lang=\"pl\"><a rel=\"mw:WikiLink\" href=\"./Montreal#Polish\" class=\"mw-selflink-fragment\">Montreal</a></span><span class=\"tpos\"><span typeof=\"mw:Entity\"> </span><a rel=\"mw:WikiLink/Interwiki\" href=\"https://pl.wiktionary.org/wiki/Montreal\" title=\"pl:Montreal\" class=\"extiw\">(pl)</a></span><span typeof=\"mw:Entity\"> </span><span class=\"gender\"><abbr title=\"masculine gender\">m</abbr><span typeof=\"mw:Entity\"> </span><abbr title=\"inanimate\">inan</abbr></span></li>\n<li>Portuguese: <span class=\"Latn\" lang=\"pt\"><a rel=\"mw:WikiLink\" href=\"./Montreal#Portuguese\" class=\"mw-selflink-fragment\">Montreal</a></span><span typeof=\"mw:Entity\"> </span><span class=\"gender\"><abbr title=\"feminine gender\">f</abbr></span>, <span class=\"Latn\" lang=\"pt\"><a rel=\"mw:WikiLink\" href=\"./Montréal#Portuguese\" title=\"Montréal\">Montréal</a></span><span typeof=\"mw:Entity\"> </span><span class=\"gender\"><abbr title=\"feminine gender\">f</abbr></span></li>\n<li>Romanian: <span class=\"Latn\" lang=\"ro\"><a rel=\"mw:WikiLink\" href=\"./Montreal#Romanian\" class=\"mw-selflink-fragment\">Montreal</a></span><span typeof=\"mw:Entity\"> </span><span class=\"gender\"><abbr title=\"neuter gender\">n</abbr></span></li>\n<li>Russian: <span class=\"Cyrl\" lang=\"ru\"><a rel=\"mw:WikiLink\" href=\"./Монреаль#Russian\" title=\"Монреаль\">Монреа́ль</a></span><span class=\"tpos\"><span typeof=\"mw:Entity\"> </span><a rel=\"mw:WikiLink/Interwiki\" href=\"https://ru.wiktionary.org/wiki/Монреаль\" title=\"ru:Монреаль\" class=\"extiw\">(ru)</a></span><span typeof=\"mw:Entity\"> </span><span class=\"gender\"><abbr title=\"masculine gender\">m</abbr></span> <span class=\"mention-gloss-paren annotation-paren\">(</span><span lang=\"ru-Latn\" class=\"tr Latn\">Monreálʹ</span><span class=\"mention-gloss-paren annotation-paren\">)</span></li>\n<li>Scottish Gaelic: <span class=\"Latn\" lang=\"gd\"><a rel=\"mw:WikiLink\" href=\"./Beinn_Rìoghail?action=edit&amp;redlink=1#Scottish_Gaelic\" title=\"Beinn Rìoghail\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Beinn Rìoghail\"]}}'>Beinn Rìoghail</a></span><span typeof=\"mw:Entity\"> </span><span class=\"gender\"><abbr title=\"feminine gender\">f</abbr></span></li>\n<li>Serbo-Croatian:\n<dl><dd>Cyrillic: <span class=\"Cyrl\" lang=\"sh\"><a rel=\"mw:WikiLink\" href=\"./Монтреал?action=edit&amp;redlink=1#Serbo-Croatian\" title=\"Монтреал\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Монтреал\"]}}'>Монтреал</a></span><span typeof=\"mw:Entity\"> </span><span class=\"gender\"><abbr title=\"masculine gender\">m</abbr></span></dd>\n<dd>Roman: <span class=\"Latn\" lang=\"sh\"><a rel=\"mw:WikiLink\" href=\"./Montreal#Serbo-Croatian\" class=\"mw-selflink-fragment\">Montreal</a></span><span class=\"tpos\"><span typeof=\"mw:Entity\"> </span><a rel=\"mw:WikiLink/Interwiki\" href=\"https://sh.wiktionary.org/wiki/Montreal\" title=\"sh:Montreal\" class=\"extiw\">(sh)</a></span><span typeof=\"mw:Entity\"> </span><span class=\"gender\"><abbr title=\"masculine gender\">m</abbr></span></dd></dl></li>\n<li>Spanish: <span class=\"Latn\" lang=\"es\"><a rel=\"mw:WikiLink\" href=\"./Montreal#Spanish\" class=\"mw-selflink-fragment\">Montreal</a></span><span class=\"tpos\"><span typeof=\"mw:Entity\"> </span><a rel=\"mw:WikiLink/Interwiki\" href=\"https://es.wiktionary.org/wiki/Montreal\" title=\"es:Montreal\" class=\"extiw\">(es)</a></span><span typeof=\"mw:Entity\"> </span><span class=\"gender\"><abbr title=\"masculine gender\">m</abbr></span></li>\n<li>Thai: <span class=\"Thai\" lang=\"th\"><a rel=\"mw:WikiLink\" href=\"./มอนทรีออล#Thai\" title=\"มอนทรีออล\">มอนทรีออล</a></span> <span class=\"mention-gloss-paren annotation-paren\">(</span><span lang=\"th-Latn\" class=\"tr Latn\">mɔn-trii-ɔ̂l</span><span class=\"mention-gloss-paren annotation-paren\">)</span></li>\n<li>Turkish: <span class=\"Latn\" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreal#Turkish\" class=\"mw-selflink-fragment\">Montreal</a></span></li>\n<li>Ukrainian: <span class=\"Cyrl\" lang=\"uk\"><a rel=\"mw:WikiLink\" href=\"./Монреаль#Ukrainian\" title=\"Монреаль\">Монреа́ль</a></span><span typeof=\"mw:Entity\"> </span><span class=\"gender\"><abbr title=\"masculine gender\">m</abbr></span> <span class=\"mention-gloss-paren annotation-paren\">(</span><span lang=\"uk-Latn\" class=\"tr Latn\">Monreálʹ</span><span class=\"mention-gloss-paren annotation-paren\">)</span></li>\n<li>Vietnamese: <span class=\"Latn\" lang=\"vi\"><a rel=\"mw:WikiLink\" href=\"./Mông-rê-an?action=edit&amp;redlink=1#Vietnamese\" title=\"Mông-rê-an\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Mông-rê-an\"]}}'>Mông-rê-an</a></span>, <span class=\"Latn\" lang=\"vi\"><a rel=\"mw:WikiLink\" href=\"./Montréal#Vietnamese\" title=\"Montréal\">Montréal</a></span></li>\n<li>Yiddish: <span typeof=\"mw:Entity\">⁧</span><span class=\"Hebr\" lang=\"yi\"><a rel=\"mw:WikiLink\" href=\"./מאָנטרעאָל#Yiddish\" title=\"מאָנטרעאָל\">מאָנטרעאָל</a></span><span typeof=\"mw:Entity\">⁩</span> <span class=\"mention-gloss-paren annotation-paren\">(</span><span lang=\"yi-Latn\" class=\"tr Latn\">montreol</span><span class=\"mention-gloss-paren annotation-paren\">)</span>, <span typeof=\"mw:Entity\">⁧</span><span class=\"Hebr\" lang=\"yi\"><a rel=\"mw:WikiLink\" href=\"./מאָנטרעאַל#Yiddish\" title=\"מאָנטרעאַל\">מאָנטרעאַל</a></span><span typeof=\"mw:Entity\">⁩</span> <span class=\"mention-gloss-paren annotation-paren\">(</span><span lang=\"yi-Latn\" class=\"tr Latn\">montreal</span><span class=\"mention-gloss-paren annotation-paren\">)</span></li></ul></td></tr>\n</tbody></table></div></div>\n\n<div class=\"NavFrame\" id=\"Translations-an_island_in_Quebec\" about=\"#mwt145\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"trans-top\",\"href\":\"./Template:trans-top\"},\"params\":{\"1\":{\"wt\":\"an island in Quebec\"}},\"i\":0}},\"\\n* French: \",{\"template\":{\"target\":{\"wt\":\"t+\",\"href\":\"./Template:t+\"},\"params\":{\"1\":{\"wt\":\"fr\"},\"2\":{\"wt\":\"Montréal\"}},\"i\":1}},\"\\n*: Middle French: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"frm\"},\"2\":{\"wt\":\"Montreal\"}},\"i\":2}},\"\\n* Laurentian: \",{\"template\":{\"target\":{\"wt\":\"t\",\"href\":\"./Template:t\"},\"params\":{\"1\":{\"wt\":\"lre\"},\"2\":{\"wt\":\"Hochelaga\"}},\"i\":3}},\"\\n\",{\"template\":{\"target\":{\"wt\":\"trans-bottom\",\"href\":\"./Template:trans-bottom\"},\"params\":{},\"i\":4}}]}'><div class=\"NavHead\" style=\"text-align:left;cursor:pointer\">an island in Quebec</div><div class=\"NavContent\"><table class=\"translations\" role=\"presentation\" data-gloss=\"an island in Quebec\" style=\"width:100%\"><tbody><tr><td class=\"translations-cell multicolumn-list\" colspan=\"3\" style=\"background-color:#ffffe0;vertical-align:top;text-align:left\">\n<ul><li>French: <span class=\"Latn\" lang=\"fr\"><a rel=\"mw:WikiLink\" href=\"./Montréal#French\" title=\"Montréal\">Montréal</a></span><span class=\"tpos\"><span typeof=\"mw:Entity\"> </span><a rel=\"mw:WikiLink/Interwiki\" href=\"https://fr.wiktionary.org/wiki/Montréal\" title=\"fr:Montréal\" class=\"extiw\">(fr)</a></span>\n<dl><dd>Middle French: <span class=\"Latn\" lang=\"frm\"><a rel=\"mw:WikiLink\" href=\"./Montreal#Middle_French\" class=\"mw-selflink-fragment\">Montreal</a></span></dd></dl></li>\n<li>Laurentian: <span class=\"Latn\" lang=\"lre\"><a rel=\"mw:WikiLink\" href=\"./Hochelaga#Laurentian\" title=\"Hochelaga\">Hochelaga</a></span></li></ul></td></tr>\n</tbody></table></div></div>\n\n</section></section><section data-mw-section-id=\"9\" id=\"mwpA\"><h3 id=\"References\">References</h3>\n<div class=\"mw-references-wrap\" typeof=\"mw:Extension/references\" about=\"#mwt151\" data-mw='{\"name\":\"references\",\"attrs\":{}}' id=\"mwpQ\"><ol class=\"mw-references references\" id=\"mwpg\"><li about=\"#cite_note-1\" id=\"cite_note-1\"><a href=\"./Montreal#cite_ref-1\" rel=\"mw:referencedBy\" id=\"mwpw\"><span class=\"mw-linkback-text\" id=\"mwqA\">↑ </span></a> <span id=\"mw-reference-text-cite_note-1\" class=\"mw-reference-text\"><span about=\"#mwt14\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"cite-web \",\"href\":\"./Template:cite-web\"},\"params\":{\"website\":{\"wt\":\"Natural Resources Canada\"},\"title\":{\"wt\":\"Geographical Name - Island of Montreal\"},\"date\":{\"wt\":\"22 August 2014\"},\"url\":{\"wt\":\"http://www.nrcan.gc.ca/earth-sciences/geography/place-names/origins-geographical-names/9218\"},\"accessdate\":{\"wt\":\"15 August 2015\"},\"archiveurl\":{\"wt\":\"https://web.archive.org/web/20160803212245/https://www.nrcan.gc.ca/earth-sciences/geography/place-names/origins-geographical-names/9218\"},\"archivedate\":{\"wt\":\"3 August 2016\"},\"author\":{\"wt\":\"Jean Poirier\"},\"publicationdate\":{\"wt\":\"1979\"}},\"i\":0}}]}' id=\"mwqQ\">Jean Poirier</span><span typeof=\"mw:Entity\" about=\"#mwt14\" id=\"mwqg\"> </span><span about=\"#mwt14\" id=\"mwqw\">(</span><span class=\"mw-formatted-date\" title=\"2014-08-22\" about=\"#mwt14\" id=\"mwrA\">22 August 2014</span><span about=\"#mwt14\" id=\"mwrQ\">),</span><span typeof=\"mw:Entity\" about=\"#mwt14\" id=\"mwrg\"> </span><span typeof=\"mw:Entity\" about=\"#mwt14\" id=\"mwrw\"> </span><span about=\"#mwt14\" id=\"mwsA\">“Geographical Name - Island of Montreal”, in</span><span typeof=\"mw:Entity\" about=\"#mwt14\" id=\"mwsQ\"> </span><span class=\"maintenance-line\" style=\"color: #777777;\" about=\"#mwt14\" id=\"mwsg\">(please provide the title of the work)</span><span typeof=\"mw:Entity\" about=\"#mwt14\" id=\"mwsw\">‎</span><sup about=\"#mwt14\" id=\"mwtA\"><a rel=\"mw:ExtLink nofollow\" href=\"https://web.archive.org/web/20160803212245/https://www.nrcan.gc.ca/earth-sciences/geography/place-names/origins-geographical-names/9218\" class=\"external autonumber\" id=\"mwtQ\"></a></sup><span about=\"#mwt14\" id=\"mwtg\">, archived from </span><a rel=\"mw:ExtLink nofollow\" href=\"https://www.nrcan.gc.ca/earth-sciences/geography/place-names/origins-geographical-names/9218\" about=\"#mwt14\" class=\"external text\" id=\"mwtw\">the original</a><span about=\"#mwt14\" id=\"mwuA\"> on </span><span class=\"mw-formatted-date\" title=\"2016-08-03\" about=\"#mwt14\" id=\"mwuQ\">2016-08-03</span></span></li></ol></div>\n\n</section><section data-mw-section-id=\"10\" id=\"mwug\"><h3 id=\"Anagrams\">Anagrams</h3>\n<ul id=\"mwuw\"><li id=\"mwvA\"><span class=\"Latn\" lang=\"en\" about=\"#mwt152\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"anagrams\",\"href\":\"./Template:anagrams\"},\"params\":{\"1\":{\"wt\":\"en\"},\"a\":{\"wt\":\"aelmnort\"},\"2\":{\"wt\":\"loan term\"},\"3\":{\"wt\":\"morantel\"}},\"i\":0}}]}' id=\"mwvQ\"><a rel=\"mw:WikiLink\" href=\"./loan_term#English\" title=\"loan term\">loan term</a></span><span about=\"#mwt152\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt152\" id=\"mwvg\"><a rel=\"mw:WikiLink\" href=\"./morantel#English\" title=\"morantel\">morantel</a></span></li></ul>\n\n</section></section><section data-mw-section-id=\"11\" id=\"mwvw\"><h2 id=\"Catalan\">Catalan</h2>\n<div class=\"sister-wikipedia sister-project noprint floatright\" style=\"border: 1px solid #aaa; font-size: 90%; background: #f9f9f9; width: 250px; padding: 4px; text-align: left;\" about=\"#mwt153\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"wikipedia\",\"href\":\"./Template:wikipedia\"},\"params\":{\"lang\":{\"wt\":\"ca\"}},\"i\":0}}]}' id=\"mwwA\"><div style=\"float: left;\"><figure class=\"mw-halign-none\" typeof=\"mw:File\"><span><img alt=\"\" resource=\"./File:Wikipedia-logo-v2.svg\" src=\"//upload.wikimedia.org/wikipedia/commons/thumb/8/80/Wikipedia-logo-v2.svg/44px-Wikipedia-logo-v2.svg.png\" decoding=\"async\" data-file-width=\"103\" data-file-height=\"94\" data-file-type=\"drawing\" height=\"40\" width=\"44\" srcset=\"//upload.wikimedia.org/wikipedia/commons/thumb/8/80/Wikipedia-logo-v2.svg/66px-Wikipedia-logo-v2.svg.png 1.5x, //upload.wikimedia.org/wikipedia/commons/thumb/8/80/Wikipedia-logo-v2.svg/88px-Wikipedia-logo-v2.svg.png 2x\" class=\"mw-file-element\"/></span><figcaption></figcaption></figure></div><div style=\"margin-left: 60px;\">Catalan <a rel=\"mw:WikiLink\" href=\"./Wikipedia\" title=\"Wikipedia\">Wikipedia</a> has an article on:<div style=\"margin-left: 10px;\"><b class=\"Latn\" lang=\"ca\"><a rel=\"mw:WikiLink/Interwiki\" href=\"https://en.wikipedia.org/wiki/ca:Montreal\" title=\"w:ca:Montreal\" class=\"extiw\">Montreal</a></b></div></div><span class=\"interProject\"><a rel=\"mw:WikiLink/Interwiki\" href=\"https://en.wikipedia.org/wiki/ca:Montreal\" title=\"w:ca:Montreal\" class=\"extiw\">Wikipedia <sup>ca</sup></a></span></div>\n\n<section data-mw-section-id=\"12\" id=\"mwwQ\"><h3 id=\"Pronunciation_2\">Pronunciation</h3>\n<ul id=\"mwwg\"><li id=\"mwww\"><a rel=\"mw:WikiLink\" href=\"./Wiktionary:International_Phonetic_Alphabet\" title=\"Wiktionary:International Phonetic Alphabet\" about=\"#mwt154\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"ca-IPA\",\"href\":\"./Template:ca-IPA\"},\"params\":{\"1\":{\"wt\":\"[Mónt.,àl]\"}},\"i\":0}}]}' id=\"mwxA\">IPA</a><sup about=\"#mwt154\">(<a rel=\"mw:WikiLink\" href=\"./Appendix:Catalan_pronunciation\" title=\"Appendix:Catalan pronunciation\">key</a>)</sup><span about=\"#mwt154\">:</span><span typeof=\"mw:Entity\" about=\"#mwt154\"> </span><span class=\"ib-brac\" about=\"#mwt154\">(</span><span class=\"ib-content\" about=\"#mwt154\"><a rel=\"mw:WikiLink/Interwiki\" href=\"https://en.wikipedia.org/wiki/Central%20Catalan\" title=\"w:Central Catalan\" class=\"extiw\">Central</a><span class=\"ib-comma\">,</span><span typeof=\"mw:Entity\"> </span><a rel=\"mw:WikiLink/Interwiki\" href=\"https://en.wikipedia.org/wiki/Balearic%20dialect\" title=\"w:Balearic dialect\" class=\"extiw\">Balearic</a><span class=\"ib-comma\">,</span><span typeof=\"mw:Entity\"> </span><a rel=\"mw:WikiLink/Interwiki\" href=\"https://en.wikipedia.org/wiki/Valencian\" title=\"w:Valencian\" class=\"extiw\">Valencian</a></span><span class=\"ib-brac\" about=\"#mwt154\">)</span><span about=\"#mwt154\"> </span><span class=\"IPA\" about=\"#mwt154\">[ˌmon.reˈal]</span><link rel=\"mw:PageProp/Category\" href=\"./Category:Catalan_terms_with_IPA_pronunciation#MONTREAL\" about=\"#mwt154\" id=\"mwxQ\"/><!-- per esAdir --></li></ul>\n\n</section><section data-mw-section-id=\"13\" id=\"mwxg\"><h3 id=\"Proper_noun_2\">Proper noun</h3>\n<p id=\"mwxw\"><span class=\"headword-line\" about=\"#mwt155\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"ca-proper noun\",\"href\":\"./Template:ca-proper_noun\"},\"params\":{\"1\":{\"wt\":\"m\"},\"head\":{\"wt\":\"[[mont|Mont]][[real]]\"}},\"i\":0}}]}' id=\"mwyA\"><strong class=\"Latn headword\" lang=\"ca\"><a rel=\"mw:WikiLink\" href=\"./mont#Catalan\" title=\"mont\">Mont</a><a rel=\"mw:WikiLink\" href=\"./real#Catalan\" title=\"real\">real</a></strong><span typeof=\"mw:Entity\"> </span><span class=\"gender\"><abbr title=\"masculine gender\">m</abbr></span></span><link rel=\"mw:PageProp/Category\" href=\"./Category:Catalan_lemmas#MONTREAL\" about=\"#mwt155\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Catalan_proper_nouns#MONTREAL\" about=\"#mwt155\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Catalan_masculine_nouns#MONTREAL\" about=\"#mwt155\" id=\"mwyQ\"/></p>\n\n<ol id=\"mwyg\"><li id=\"mwyw\"><span class=\"Latn\" lang=\"en\" about=\"#mwt156\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"place\",\"href\":\"./Template:place\"},\"params\":{\"1\":{\"wt\":\"ca\"},\"2\":{\"wt\":\"largest city\"},\"3\":{\"wt\":\"p/Quebec\"},\"4\":{\"wt\":\"c/Canada\"},\"t1\":{\"wt\":\"Montreal\"}},\"i\":0}}]}' id=\"mwzA\"><a rel=\"mw:WikiLink\" href=\"./Montreal#English\" class=\"mw-selflink-fragment\">Montreal</a></span><span about=\"#mwt156\"> (the largest city in </span><span class=\"Latn\" lang=\"en\" about=\"#mwt156\"><a rel=\"mw:WikiLink\" href=\"./Quebec#English\" title=\"Quebec\">Quebec</a></span><span about=\"#mwt156\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt156\"><a rel=\"mw:WikiLink\" href=\"./Canada#English\" title=\"Canada\">Canada</a></span><span about=\"#mwt156\">)</span><link rel=\"mw:PageProp/Category\" href=\"./Category:ca:Montreal#MONTREAL\" about=\"#mwt156\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:ca:Montreal#MONTREAL\" about=\"#mwt156\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:ca:Cities_in_Quebec#MONTREAL\" about=\"#mwt156\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:ca:Cities_in_Canada#MONTREAL\" about=\"#mwt156\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:ca:Places_in_Quebec#MONTREAL\" about=\"#mwt156\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:ca:Places_in_Canada#MONTREAL\" about=\"#mwt156\" id=\"mwzQ\"/></li>\n<li id=\"mwzg\"><span class=\"Latn\" lang=\"en\" about=\"#mwt157\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"place\",\"href\":\"./Template:place\"},\"params\":{\"1\":{\"wt\":\"ca\"},\"2\":{\"wt\":\"island\"},\"3\":{\"wt\":\"p/Quebec\"},\"4\":{\"wt\":\"c/Canada\"},\"5\":{\"wt\":\"the site of the city of [[Montreal]]\"},\"t1\":{\"wt\":\"Montreal\"}},\"i\":0}}]}' id=\"mwzw\"><a rel=\"mw:WikiLink\" href=\"./Montreal#English\" class=\"mw-selflink-fragment\">Montreal</a></span><span about=\"#mwt157\"> (an island in </span><span class=\"Latn\" lang=\"en\" about=\"#mwt157\"><a rel=\"mw:WikiLink\" href=\"./Quebec#English\" title=\"Quebec\">Quebec</a></span><span about=\"#mwt157\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt157\"><a rel=\"mw:WikiLink\" href=\"./Canada#English\" title=\"Canada\">Canada</a></span><span about=\"#mwt157\">, the site of the city of </span><a rel=\"mw:WikiLink\" href=\"./Montreal\" about=\"#mwt157\" class=\"mw-selflink selflink\">Montreal</a><span about=\"#mwt157\">)</span><link rel=\"mw:PageProp/Category\" href=\"./Category:ca:Islands#MONTREAL\" about=\"#mwt157\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:ca:Places_in_Quebec#MONTREAL\" about=\"#mwt157\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:ca:Places_in_Canada#MONTREAL\" about=\"#mwt157\" id=\"mw0A\"/></li></ol>\n\n</section></section><section data-mw-section-id=\"14\" id=\"mw0Q\"><h2 id=\"German\">German</h2>\n<div class=\"sister-wikipedia sister-project noprint floatright\" style=\"border: 1px solid #aaa; font-size: 90%; background: #f9f9f9; width: 250px; padding: 4px; text-align: left;\" about=\"#mwt158\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"wikipedia\",\"href\":\"./Template:wikipedia\"},\"params\":{\"lang\":{\"wt\":\"de\"}},\"i\":0}}]}' id=\"mw0g\"><div style=\"float: left;\"><figure class=\"mw-halign-none\" typeof=\"mw:File\"><span><img alt=\"\" resource=\"./File:Wikipedia-logo-v2.svg\" src=\"//upload.wikimedia.org/wikipedia/commons/thumb/8/80/Wikipedia-logo-v2.svg/44px-Wikipedia-logo-v2.svg.png\" decoding=\"async\" data-file-width=\"103\" data-file-height=\"94\" data-file-type=\"drawing\" height=\"40\" width=\"44\" srcset=\"//upload.wikimedia.org/wikipedia/commons/thumb/8/80/Wikipedia-logo-v2.svg/66px-Wikipedia-logo-v2.svg.png 1.5x, //upload.wikimedia.org/wikipedia/commons/thumb/8/80/Wikipedia-logo-v2.svg/88px-Wikipedia-logo-v2.svg.png 2x\" class=\"mw-file-element\"/></span><figcaption></figcaption></figure></div><div style=\"margin-left: 60px;\">German <a rel=\"mw:WikiLink\" href=\"./Wikipedia\" title=\"Wikipedia\">Wikipedia</a> has an article on:<div style=\"margin-left: 10px;\"><b class=\"Latn\" lang=\"de\"><a rel=\"mw:WikiLink/Interwiki\" href=\"https://en.wikipedia.org/wiki/de:Montreal\" title=\"w:de:Montreal\" class=\"extiw\">Montreal</a></b></div></div><span class=\"interProject\"><a rel=\"mw:WikiLink/Interwiki\" href=\"https://en.wikipedia.org/wiki/de:Montreal\" title=\"w:de:Montreal\" class=\"extiw\">Wikipedia <sup>de</sup></a></span></div>\n\n<section data-mw-section-id=\"15\" id=\"mw0w\"><h3 id=\"Pronunciation_3\">Pronunciation</h3>\n<ul id=\"mw1A\"><li id=\"mw1Q\"><a rel=\"mw:WikiLink\" href=\"./Wiktionary:International_Phonetic_Alphabet\" title=\"Wiktionary:International Phonetic Alphabet\" about=\"#mwt159\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"IPA\",\"href\":\"./Template:IPA\"},\"params\":{\"1\":{\"wt\":\"de\"},\"2\":{\"wt\":\"[mɔntʁeˈaːl]\"}},\"i\":0}}]}' id=\"mw1g\">IPA</a><sup about=\"#mwt159\">(<a rel=\"mw:WikiLink\" href=\"./Appendix:German_pronunciation\" title=\"Appendix:German pronunciation\">key</a>)</sup><span about=\"#mwt159\">:</span><span typeof=\"mw:Entity\" about=\"#mwt159\"> </span><span class=\"IPA\" about=\"#mwt159\">[mɔntʁeˈaːl]</span><link rel=\"mw:PageProp/Category\" href=\"./Category:German_terms_with_IPA_pronunciation#MONTREAL\" about=\"#mwt159\" id=\"mw1w\"/></li>\n<li id=\"mw2A\"><link rel=\"mw-deduplicated-inline-style\" href=\"mw-data:TemplateStyles:r50165410\" about=\"#mwt160\" typeof=\"mw:Extension/templatestyles mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"audio\",\"href\":\"./Template:audio\"},\"params\":{\"1\":{\"wt\":\"de\"},\"2\":{\"wt\":\"De-Montreal.ogg\"},\"3\":{\"wt\":\"Audio\"}},\"i\":0}}]}' id=\"mw2Q\"/><table class=\"audiotable\" style=\"vertical-align: middle; display:inline-block; list-style:none;line-height: 1em; border-collapse:collapse;\" about=\"#mwt160\"><tbody><tr><td class=\"unicode audiolink\" style=\"padding-right:5px; padding-left: 0;\">Audio</td><td class=\"audiofile\"><span typeof=\"mw:File\" class=\"mw-default-audio-height\" data-mw='{\"caption\":\"noicon\"}'><span title=\"noicon\"><audio controls=\"\" preload=\"none\" height=\"32\" width=\"175\" resource=\"./File:De-Montreal.ogg\" data-durationhint=\"2\" class=\"mw-file-element\"><source src=\"//upload.wikimedia.org/wikipedia/commons/b/b9/De-Montreal.ogg\" type='audio/ogg; codecs=\"vorbis\"'/><source src=\"//upload.wikimedia.org/wikipedia/commons/transcoded/b/b9/De-Montreal.ogg/De-Montreal.ogg.mp3\" type=\"audio/mpeg\" data-transcodekey=\"mp3\"/></audio></span></span></td><td class=\"audiometa\" style=\"font-size: 80%;\">(<a rel=\"mw:WikiLink\" href=\"./File:De-Montreal.ogg\" title=\"File:De-Montreal.ogg\">file</a>)</td></tr></tbody></table><link rel=\"mw:PageProp/Category\" href=\"./Category:German_terms_with_audio_links#MONTREAL\" about=\"#mwt160\" id=\"mw2g\"/></li></ul>\n\n</section><section data-mw-section-id=\"16\" id=\"mw2w\"><h3 id=\"Proper_noun_3\">Proper noun</h3>\n<p id=\"mw3A\"><span class=\"headword-line\" about=\"#mwt163\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"de-proper noun\",\"href\":\"./Template:de-proper_noun\"},\"params\":{\"1\":{\"wt\":\"toponym\"}},\"i\":0}}]}' id=\"mw3Q\"><strong class=\"Latn headword\" lang=\"de\">Montreal</strong><span typeof=\"mw:Entity\"> </span><span class=\"gender\"><abbr title=\"neuter gender\">n</abbr></span> (<i><a rel=\"mw:WikiLink\" href=\"./Appendix:Glossary#proper_noun\" title=\"Appendix:Glossary\">proper noun</a></i>, <i>genitive</i> <b class=\"Latn\" lang=\"de\"><a rel=\"mw:WikiLink\" href=\"./Montreals?action=edit&amp;redlink=1#German\" title=\"Montreals\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreals\"]}}'>Montreals</a></b> <i>or</i> <span class=\"ib-brac qualifier-brac\">(</span><span class=\"ib-content qualifier-content\">optionally with an article</span><span class=\"ib-brac qualifier-brac\">)</span> <b class=\"Latn\" lang=\"de\"><a rel=\"mw:WikiLink\" href=\"./Montreal#German\" class=\"mw-selflink-fragment\">Montreal</a></b>)</span><link rel=\"mw:PageProp/Category\" href=\"./Category:German_lemmas#MONTREAL\" about=\"#mwt163\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:German_proper_nouns#MONTREAL\" about=\"#mwt163\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:German_neuter_nouns#MONTREAL\" about=\"#mwt163\" id=\"mw3g\"/></p>\n\n<ol id=\"mw3w\"><li id=\"mw4A\"><span class=\"Latn\" lang=\"en\" about=\"#mwt164\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"place\",\"href\":\"./Template:place\"},\"params\":{\"1\":{\"wt\":\"de\"},\"2\":{\"wt\":\"largest city\"},\"3\":{\"wt\":\"p/Quebec\"},\"4\":{\"wt\":\"c/Canada\"},\"t1\":{\"wt\":\"Montreal\"}},\"i\":0}}]}' id=\"mw4Q\"><a rel=\"mw:WikiLink\" href=\"./Montreal#English\" class=\"mw-selflink-fragment\">Montreal</a></span><span about=\"#mwt164\"> (the largest city in </span><span class=\"Latn\" lang=\"en\" about=\"#mwt164\"><a rel=\"mw:WikiLink\" href=\"./Quebec#English\" title=\"Quebec\">Quebec</a></span><span about=\"#mwt164\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt164\"><a rel=\"mw:WikiLink\" href=\"./Canada#English\" title=\"Canada\">Canada</a></span><span about=\"#mwt164\">)</span><link rel=\"mw:PageProp/Category\" href=\"./Category:de:Montreal#MONTREAL\" about=\"#mwt164\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:de:Montreal#MONTREAL\" about=\"#mwt164\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:de:Cities_in_Quebec#MONTREAL\" about=\"#mwt164\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:de:Cities_in_Canada#MONTREAL\" about=\"#mwt164\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:de:Places_in_Quebec#MONTREAL\" about=\"#mwt164\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:de:Places_in_Canada#MONTREAL\" about=\"#mwt164\" id=\"mw4g\"/></li></ol>\n\n</section></section><section data-mw-section-id=\"17\" id=\"mw4w\"><h2 id=\"Middle_French\">Middle French</h2>\n\n<section data-mw-section-id=\"18\" id=\"mw5A\"><h3 id=\"Etymology_2\">Etymology</h3>\n<p id=\"mw5Q\"><i class=\"Latn mention\" lang=\"frm\" about=\"#mwt165\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"compound\",\"href\":\"./Template:compound\"},\"params\":{\"1\":{\"wt\":\"frm\"},\"2\":{\"wt\":\"mont\"},\"3\":{\"wt\":\"real\"}},\"i\":0}}]}' id=\"mw5g\"><a rel=\"mw:WikiLink\" href=\"./mont#Middle_French\" title=\"mont\">mont</a></i><span about=\"#mwt165\"> +</span><span typeof=\"mw:Entity\" about=\"#mwt165\">‎</span><span about=\"#mwt165\"> </span><i class=\"Latn mention\" lang=\"frm\" about=\"#mwt165\"><a rel=\"mw:WikiLink\" href=\"./real#Middle_French\" title=\"real\">real</a></i><link rel=\"mw:PageProp/Category\" href=\"./Category:Middle_French_compound_terms#MONTREAL\" about=\"#mwt165\" id=\"mw5w\"/>; The island was named after the mountain located on it, <i class=\"Latn mention\" lang=\"frm\" about=\"#mwt166\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"m\",\"href\":\"./Template:m\"},\"params\":{\"1\":{\"wt\":\"frm\"},\"2\":{\"wt\":\"mont Royal\"}},\"i\":0}}]}' id=\"mw6A\"><a rel=\"mw:WikiLink\" href=\"./mont_Royal?action=edit&amp;redlink=1#Middle_French\" title=\"mont Royal\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"mont Royal\"]}}'>mont Royal</a></i>, which was named by Jacques Cartier.</p>\n\n</section><section data-mw-section-id=\"19\" id=\"mw6Q\"><h3 id=\"Proper_noun_4\">Proper noun</h3>\n<p id=\"mw6g\"><span class=\"headword-line\" about=\"#mwt167\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"head\",\"href\":\"./Template:head\"},\"params\":{\"1\":{\"wt\":\"frm\"},\"2\":{\"wt\":\"proper noun\"}},\"i\":0}}]}' id=\"mw6w\"><strong class=\"Latn headword\" lang=\"frm\">Montreal</strong></span><link rel=\"mw:PageProp/Category\" href=\"./Category:Middle_French_lemmas#MONTREAL\" about=\"#mwt167\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Middle_French_proper_nouns#MONTREAL\" about=\"#mwt167\" id=\"mw7A\"/></p>\n\n<ol id=\"mw7Q\"><li id=\"mw7g\"><span about=\"#mwt168\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"place\",\"href\":\"./Template:place\"},\"params\":{\"1\":{\"wt\":\"frm\"},\"2\":{\"wt\":\"island\"},\"3\":{\"wt\":\"riv:Suf/Saint Lawrence\"},\"4\":{\"wt\":\"in modern\"},\"5\":{\"wt\":\"c/Canada\"}},\"i\":0}}]}' id=\"mw7w\">An island in the </span><span class=\"Latn\" lang=\"en\" about=\"#mwt168\"><a rel=\"mw:WikiLink\" href=\"./Saint_Lawrence?action=edit&amp;redlink=1#English\" title=\"Saint Lawrence\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Saint Lawrence\"]}}'>Saint Lawrence</a></span><span about=\"#mwt168\"> River, in modern </span><span class=\"Latn\" lang=\"en\" about=\"#mwt168\"><a rel=\"mw:WikiLink\" href=\"./Canada#English\" title=\"Canada\">Canada</a></span><link rel=\"mw:PageProp/Category\" href=\"./Category:frm:Islands#MONTREAL\" about=\"#mwt168\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:frm:Places_in_Canada#MONTREAL\" about=\"#mwt168\" id=\"mw8A\"/>, at its confluence with the <a rel=\"mw:WikiLink\" href=\"./Ottawa\" title=\"Ottawa\" id=\"mw8Q\">Ottawa</a> River.</li></ol>\n\n<section data-mw-section-id=\"20\" id=\"mw8g\"><h4 id=\"Descendants\">Descendants</h4>\n<div class=\"derivedterms ul-column-count\" data-column-count=\"2\" about=\"#mwt169\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"top2\",\"href\":\"./Template:top2\"},\"params\":{},\"i\":0}},\"\\n* \",{\"template\":{\"target\":{\"wt\":\"desc\",\"href\":\"./Template:desc\"},\"params\":{\"1\":{\"wt\":\"en\"},\"2\":{\"wt\":\"Montreal\"}},\"i\":1}},\"\\n* \",{\"template\":{\"target\":{\"wt\":\"desc\",\"href\":\"./Template:desc\"},\"params\":{\"1\":{\"wt\":\"fr\"},\"2\":{\"wt\":\"Montréal\"}},\"i\":2}},\"\\n\",{\"template\":{\"target\":{\"wt\":\"bottom\",\"href\":\"./Template:bottom\"},\"params\":{},\"i\":3}}]}' id=\"mw8w\">\n<ul><li>English: <span class=\"Latn\" lang=\"en\"><a rel=\"mw:WikiLink\" href=\"./Montreal#English\" class=\"mw-selflink-fragment\">Montreal</a></span><style data-mw-deduplicate=\"TemplateStyles:r68481116\" typeof=\"mw:Extension/templatestyles\" about=\"#mwt172\" data-mw='{\"name\":\"templatestyles\",\"attrs\":{\"src\":\"Module:etymology/style.css\"},\"body\":{\"extsrc\":\"\"}}'>.mw-parser-output .desc-arr[title]{cursor:help}.mw-parser-output .desc-arr[title=\"uncertain\"]{font-size:.7em;vertical-align:super}</style></li>\n<li>French: <span class=\"Latn\" lang=\"fr\"><a rel=\"mw:WikiLink\" href=\"./Montréal#French\" title=\"Montréal\">Montréal</a></span><link rel=\"mw-deduplicated-inline-style\" href=\"mw-data:TemplateStyles:r68481116\" about=\"#mwt175\" typeof=\"mw:Extension/templatestyles\" data-mw='{\"name\":\"templatestyles\",\"attrs\":{\"src\":\"Module:etymology/style.css\"},\"body\":{\"extsrc\":\"\"}}'/></li></ul>\n\n</div>\n\n</section></section></section><section data-mw-section-id=\"21\" id=\"mw9A\"><h2 id=\"Norwegian_Bokmål\"><span id=\"Norwegian_Bokm.C3.A5l\" typeof=\"mw:FallbackId\"></span>Norwegian Bokmål</h2>\n\n<section data-mw-section-id=\"22\" id=\"mw9Q\"><h3 id=\"Proper_noun_5\">Proper noun</h3>\n<p id=\"mw9g\"><span class=\"headword-line\" about=\"#mwt177\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"head\",\"href\":\"./Template:head\"},\"params\":{\"1\":{\"wt\":\"nb\"},\"2\":{\"wt\":\"proper noun\"}},\"i\":0}}]}' id=\"mw9w\"><strong class=\"Latn headword\" lang=\"nb\">Montreal</strong></span><link rel=\"mw:PageProp/Category\" href=\"./Category:Norwegian_Bokmål_lemmas#MONTREAL\" about=\"#mwt177\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Norwegian_Bokmål_proper_nouns#MONTREAL\" about=\"#mwt177\" id=\"mw-A\"/></p>\n\n<ol id=\"mw-Q\"><li id=\"mw-g\"><span class=\"form-of-definition use-with-mention\" about=\"#mwt178\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"alt form\",\"href\":\"./Template:alt_form\"},\"params\":{\"1\":{\"wt\":\"nb\"},\"2\":{\"wt\":\"Montréal\"},\"nocap\":{\"wt\":\"1\"}},\"i\":0}}]}' id=\"mw-w\">alternative form of <span class=\"form-of-definition-link\"><i class=\"Latn mention\" lang=\"nb\"><a rel=\"mw:WikiLink\" href=\"./Montréal#Norwegian_Bokmål\" title=\"Montréal\">Montréal</a></i></span></span>, <span about=\"#mwt179\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"place\",\"href\":\"./Template:place\"},\"params\":{\"1\":{\"wt\":\"nb\"},\"a\":{\"wt\":\"the\"},\"2\":{\"wt\":\"largest city\"},\"3\":{\"wt\":\"p/Quebec\"},\"4\":{\"wt\":\"c/Canada\"}},\"i\":0}}]}' id=\"mw_A\">the largest city in </span><span class=\"Latn\" lang=\"en\" about=\"#mwt179\"><a rel=\"mw:WikiLink\" href=\"./Quebec#English\" title=\"Quebec\">Quebec</a></span><span about=\"#mwt179\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt179\"><a rel=\"mw:WikiLink\" href=\"./Canada#English\" title=\"Canada\">Canada</a></span><link rel=\"mw:PageProp/Category\" href=\"./Category:nb:Montreal#MONTREAL\" about=\"#mwt179\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:nb:Cities_in_Quebec#MONTREAL\" about=\"#mwt179\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:nb:Cities_in_Canada#MONTREAL\" about=\"#mwt179\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:nb:Places_in_Quebec#MONTREAL\" about=\"#mwt179\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:nb:Places_in_Canada#MONTREAL\" about=\"#mwt179\" id=\"mw_Q\"/></li></ol>\n\n</section></section><section data-mw-section-id=\"23\" id=\"mw_g\"><h2 id=\"Norwegian_Nynorsk\">Norwegian Nynorsk</h2>\n\n<section data-mw-section-id=\"24\" id=\"mw_w\"><h3 id=\"Proper_noun_6\">Proper noun</h3>\n<p id=\"mwAQA\"><span class=\"headword-line\" about=\"#mwt180\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"head\",\"href\":\"./Template:head\"},\"params\":{\"1\":{\"wt\":\"nn\"},\"2\":{\"wt\":\"proper noun\"}},\"i\":0}}]}' id=\"mwAQE\"><strong class=\"Latn headword\" lang=\"nn\">Montreal</strong></span><link rel=\"mw:PageProp/Category\" href=\"./Category:Norwegian_Nynorsk_lemmas#MONTREAL\" about=\"#mwt180\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Norwegian_Nynorsk_proper_nouns#MONTREAL\" about=\"#mwt180\" id=\"mwAQI\"/></p>\n\n<ol id=\"mwAQM\"><li id=\"mwAQQ\"><span class=\"form-of-definition use-with-mention\" about=\"#mwt181\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"alt form\",\"href\":\"./Template:alt_form\"},\"params\":{\"1\":{\"wt\":\"nn\"},\"2\":{\"wt\":\"Montréal\"},\"nocap\":{\"wt\":\"1\"}},\"i\":0}}]}' id=\"mwAQU\">alternative form of <span class=\"form-of-definition-link\"><i class=\"Latn mention\" lang=\"nn\"><a rel=\"mw:WikiLink\" href=\"./Montreal#Norwegian_Nynorsk\" class=\"mw-selflink-fragment\">Montréal</a></i></span></span><link rel=\"mw:PageProp/Category\" href=\"./Category:Forms_linking_to_themselves#MONTREAL\" about=\"#mwt181\" id=\"mwAQY\"/>, <span about=\"#mwt182\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"place\",\"href\":\"./Template:place\"},\"params\":{\"1\":{\"wt\":\"nn\"},\"a\":{\"wt\":\"the\"},\"2\":{\"wt\":\"largest city\"},\"3\":{\"wt\":\"p/Quebec\"},\"4\":{\"wt\":\"c/Canada\"}},\"i\":0}}]}' id=\"mwAQc\">the largest city in </span><span class=\"Latn\" lang=\"en\" about=\"#mwt182\"><a rel=\"mw:WikiLink\" href=\"./Quebec#English\" title=\"Quebec\">Quebec</a></span><span about=\"#mwt182\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt182\"><a rel=\"mw:WikiLink\" href=\"./Canada#English\" title=\"Canada\">Canada</a></span><link rel=\"mw:PageProp/Category\" href=\"./Category:nn:Montreal#MONTREAL\" about=\"#mwt182\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:nn:Cities_in_Quebec#MONTREAL\" about=\"#mwt182\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:nn:Cities_in_Canada#MONTREAL\" about=\"#mwt182\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:nn:Places_in_Quebec#MONTREAL\" about=\"#mwt182\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:nn:Places_in_Canada#MONTREAL\" about=\"#mwt182\" id=\"mwAQg\"/></li></ol>\n\n</section></section><section data-mw-section-id=\"25\" id=\"mwAQk\"><h2 id=\"Polish\">Polish</h2>\n<div class=\"sister-wikipedia sister-project noprint floatright\" style=\"border: 1px solid #aaa; font-size: 90%; background: #f9f9f9; width: 250px; padding: 4px; text-align: left;\" about=\"#mwt183\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"wp\",\"href\":\"./Template:wp\"},\"params\":{\"lang\":{\"wt\":\"pl\"}},\"i\":0}}]}' id=\"mwAQo\"><div style=\"float: left;\"><figure class=\"mw-halign-none\" typeof=\"mw:File\"><span><img alt=\"\" resource=\"./File:Wikipedia-logo-v2.svg\" src=\"//upload.wikimedia.org/wikipedia/commons/thumb/8/80/Wikipedia-logo-v2.svg/44px-Wikipedia-logo-v2.svg.png\" decoding=\"async\" data-file-width=\"103\" data-file-height=\"94\" data-file-type=\"drawing\" height=\"40\" width=\"44\" srcset=\"//upload.wikimedia.org/wikipedia/commons/thumb/8/80/Wikipedia-logo-v2.svg/66px-Wikipedia-logo-v2.svg.png 1.5x, //upload.wikimedia.org/wikipedia/commons/thumb/8/80/Wikipedia-logo-v2.svg/88px-Wikipedia-logo-v2.svg.png 2x\" class=\"mw-file-element\"/></span><figcaption></figcaption></figure></div><div style=\"margin-left: 60px;\">Polish <a rel=\"mw:WikiLink\" href=\"./Wikipedia\" title=\"Wikipedia\">Wikipedia</a> has an article on:<div style=\"margin-left: 10px;\"><b class=\"Latn\" lang=\"pl\"><a rel=\"mw:WikiLink/Interwiki\" href=\"https://en.wikipedia.org/wiki/pl:Montreal\" title=\"w:pl:Montreal\" class=\"extiw\">Montreal</a></b></div></div><span class=\"interProject\"><a rel=\"mw:WikiLink/Interwiki\" href=\"https://en.wikipedia.org/wiki/pl:Montreal\" title=\"w:pl:Montreal\" class=\"extiw\">Wikipedia <sup>pl</sup></a></span></div>\n\n<section data-mw-section-id=\"26\" id=\"mwAQs\"><h3 id=\"Etymology_3\">Etymology</h3>\n<p id=\"mwAQw\"><a rel=\"mw:WikiLink\" href=\"./Appendix:Glossary#loanword\" title=\"Appendix:Glossary\" about=\"#mwt184\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"bor+\",\"href\":\"./Template:bor+\"},\"params\":{\"1\":{\"wt\":\"pl\"},\"2\":{\"wt\":\"en\"},\"3\":{\"wt\":\"Montreal\"}},\"i\":0}}]}' id=\"mwAQ0\">Borrowed</a><span about=\"#mwt184\"> from </span><span class=\"etyl\" about=\"#mwt184\"><a rel=\"mw:WikiLink/Interwiki\" href=\"https://en.wikipedia.org/wiki/English%20language\" title=\"w:English language\" class=\"extiw\">English</a><link rel=\"mw:PageProp/Category\" href=\"./Category:Polish_terms_borrowed_from_English#MONTREAL\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Polish_terms_derived_from_English#MONTREAL\"/></span><span about=\"#mwt184\"> </span><i class=\"Latn mention\" lang=\"en\" about=\"#mwt184\" id=\"mwAQ4\"><a rel=\"mw:WikiLink\" href=\"./Montreal#English\" class=\"mw-selflink-fragment\">Montreal</a></i>, from <span class=\"etyl\" about=\"#mwt185\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"der\",\"href\":\"./Template:der\"},\"params\":{\"1\":{\"wt\":\"pl\"},\"2\":{\"wt\":\"fr\"},\"3\":{\"wt\":\"Montréal\"}},\"i\":0}}]}' id=\"mwAQ8\"><a rel=\"mw:WikiLink/Interwiki\" href=\"https://en.wikipedia.org/wiki/French%20language\" title=\"w:French language\" class=\"extiw\">French</a><link rel=\"mw:PageProp/Category\" href=\"./Category:Polish_terms_derived_from_French#MONTREAL\"/></span><span about=\"#mwt185\"> </span><i class=\"Latn mention\" lang=\"fr\" about=\"#mwt185\" id=\"mwARA\"><a rel=\"mw:WikiLink\" href=\"./Montréal#French\" title=\"Montréal\">Montréal</a></i>, from <i class=\"Latn mention\" lang=\"fr\" about=\"#mwt186\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"m\",\"href\":\"./Template:m\"},\"params\":{\"1\":{\"wt\":\"fr\"},\"2\":{\"wt\":\"Mont Royal\"},\"gloss\":{\"wt\":\"[[Mount Royal]]\"}},\"i\":0}}]}' id=\"mwARE\"><a rel=\"mw:WikiLink\" href=\"./Mont_Royal?action=edit&amp;redlink=1#French\" title=\"Mont Royal\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Mont Royal\"]}}'>Mont Royal</a></i><span about=\"#mwt186\"> </span><span class=\"mention-gloss-paren annotation-paren\" about=\"#mwt186\">(</span><span class=\"mention-gloss-double-quote\" about=\"#mwt186\">“</span><span class=\"mention-gloss\" about=\"#mwt186\"><a rel=\"mw:WikiLink\" href=\"./Mount_Royal\" title=\"Mount Royal\">Mount Royal</a></span><span class=\"mention-gloss-double-quote\" about=\"#mwt186\">”</span><span class=\"mention-gloss-paren annotation-paren\" about=\"#mwt186\" id=\"mwARI\">)</span>, after the nearby mountain by Jacques Cartier.</p>\n\n</section><section data-mw-section-id=\"27\" id=\"mwARM\"><h3 id=\"Pronunciation_4\">Pronunciation</h3>\n<ul about=\"#mwt187\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"pl-p\",\"href\":\"./Template:pl-p\"},\"params\":{\"a\":{\"wt\":\"Pl-Montreal.ogg\"}},\"i\":0}}]}' id=\"mwARQ\"><li><a rel=\"mw:WikiLink\" href=\"./Wiktionary:International_Phonetic_Alphabet\" title=\"Wiktionary:International Phonetic Alphabet\">IPA</a><sup>(<a rel=\"mw:WikiLink\" href=\"./Appendix:Polish_pronunciation\" title=\"Appendix:Polish pronunciation\">key</a>)</sup>:<span typeof=\"mw:Entity\"> </span><span class=\"IPA\">/mɔnˈtrɛ.al/</span><link rel=\"mw:PageProp/Category\" href=\"./Category:Polish_3-syllable_words\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Polish_terms_with_IPA_pronunciation#MONTREAL\"/></li>\n<li><link rel=\"mw-deduplicated-inline-style\" href=\"mw-data:TemplateStyles:r50165410\" about=\"#mwt189\" typeof=\"mw:Extension/templatestyles\" data-mw='{\"name\":\"templatestyles\",\"attrs\":{\"src\":\"audio/styles.css\"}}'/><table class=\"audiotable\" style=\"vertical-align: middle; display:inline-block; list-style:none;line-height: 1em; border-collapse:collapse;\"><tbody><tr><td class=\"unicode audiolink\" style=\"padding-right:5px; padding-left: 0;\">Audio</td><td class=\"audiofile\"><span typeof=\"mw:File\" class=\"mw-default-audio-height\" data-mw='{\"caption\":\"noicon\"}'><span title=\"noicon\"><audio controls=\"\" preload=\"none\" height=\"32\" width=\"175\" resource=\"./File:Pl-Montreal.ogg\" data-durationhint=\"2\" class=\"mw-file-element\"><source src=\"//upload.wikimedia.org/wikipedia/commons/0/01/Pl-Montreal.ogg\" type='audio/ogg; codecs=\"vorbis\"'/><source src=\"//upload.wikimedia.org/wikipedia/commons/transcoded/0/01/Pl-Montreal.ogg/Pl-Montreal.ogg.mp3\" type=\"audio/mpeg\" data-transcodekey=\"mp3\"/></audio></span></span></td><td class=\"audiometa\" style=\"font-size: 80%;\">(<a rel=\"mw:WikiLink\" href=\"./File:Pl-Montreal.ogg\" title=\"File:Pl-Montreal.ogg\">file</a>)</td></tr></tbody></table><link rel=\"mw:PageProp/Category\" href=\"./Category:Polish_terms_with_audio_links#MONTREAL\"/></li>\n<li>Rhymes: <a rel=\"mw:WikiLink\" href=\"./Rhymes:Polish/ɛal?action=edit&amp;redlink=1\" title=\"Rhymes:Polish/ɛal\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Rhymes:Polish/ɛal\"]}}'><span class=\"IPA\">-ɛal</span></a><link rel=\"mw:PageProp/Category\" href=\"./Category:Rhymes:Polish/ɛal#MONTREAL\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Rhymes:Polish/ɛal/3_syllables#MONTREAL\"/></li>\n<li>Syllabification: <span class=\"Latn\" lang=\"pl\">Mon‧tre‧al</span></li></ul>\n\n</section><section data-mw-section-id=\"28\" id=\"mwARU\"><h3 id=\"Proper_noun_7\">Proper noun</h3>\n<p id=\"mwARY\"><span class=\"headword-line\" about=\"#mwt190\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"pl-proper noun\",\"href\":\"./Template:pl-proper_noun\"},\"params\":{\"1\":{\"wt\":\"m-in\"}},\"i\":0}}]}' id=\"mwARc\"><strong class=\"Latn headword\" lang=\"pl\">Montreal</strong><span typeof=\"mw:Entity\"> </span><span class=\"gender\"><abbr title=\"masculine gender\">m</abbr><span typeof=\"mw:Entity\"> </span><abbr title=\"inanimate\">inan</abbr></span></span><link rel=\"mw:PageProp/Category\" href=\"./Category:Polish_lemmas#MONTREAL\" about=\"#mwt190\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Polish_proper_nouns#MONTREAL\" about=\"#mwt190\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Polish_masculine_nouns#MONTREAL\" about=\"#mwt190\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Polish_inanimate_nouns#MONTREAL\" about=\"#mwt190\" id=\"mwARg\"/></p>\n\n<ol id=\"mwARk\"><li id=\"mwARo\"><span class=\"Latn\" lang=\"en\" about=\"#mwt191\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"place\",\"href\":\"./Template:place\"},\"params\":{\"1\":{\"wt\":\"pl\"},\"2\":{\"wt\":\"largest city\"},\"3\":{\"wt\":\"p/Quebec\"},\"4\":{\"wt\":\"c/Canada\"},\"t1\":{\"wt\":\"Montreal\"}},\"i\":0}}]}' id=\"mwARs\"><a rel=\"mw:WikiLink\" href=\"./Montreal#English\" class=\"mw-selflink-fragment\">Montreal</a></span><span about=\"#mwt191\"> (the largest city in </span><span class=\"Latn\" lang=\"en\" about=\"#mwt191\"><a rel=\"mw:WikiLink\" href=\"./Quebec#English\" title=\"Quebec\">Quebec</a></span><span about=\"#mwt191\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt191\"><a rel=\"mw:WikiLink\" href=\"./Canada#English\" title=\"Canada\">Canada</a></span><span about=\"#mwt191\">)</span><link rel=\"mw:PageProp/Category\" href=\"./Category:pl:Montreal#MONTREAL\" about=\"#mwt191\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:pl:Montreal#MONTREAL\" about=\"#mwt191\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:pl:Cities_in_Quebec#MONTREAL\" about=\"#mwt191\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:pl:Cities_in_Canada#MONTREAL\" about=\"#mwt191\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:pl:Places_in_Quebec#MONTREAL\" about=\"#mwt191\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:pl:Places_in_Canada#MONTREAL\" about=\"#mwt191\" id=\"mwARw\"/></li></ol>\n\n<section data-mw-section-id=\"29\" id=\"mwAR0\"><h4 id=\"Declension\">Declension</h4>\n<div class=\"NavFrame\" style=\"display: block; max-width: 19em;\" about=\"#mwt192\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"pl-decl-noun-m-in\",\"href\":\"./Template:pl-decl-noun-m-in\"},\"params\":{\"tantum\":{\"wt\":\"s\"}},\"i\":0}}]}' id=\"mwAR4\">\n<div class=\"NavHead\" style=\"background:#eff7ff\">Declension of <span class=\"Latn mention\" lang=\"pl\">Montreal</span></div>\n<div class=\"NavContent\">\n<table style=\"background:#F9F9F9; text-align:center; width: 100%; margin: 0;\" class=\"inflection-table\">\n<tbody><tr>\n<th style=\"background:#d9ebff; width: 8em;\"></th>\n<th style=\"background:#d9ebff;\" scope=\"col\">singular</th></tr>\n<tr>\n<th title=\"mianownik (kto? co?)\" style=\"background:#eff7ff;\" scope=\"row\">nominative</th>\n<td><span class=\"Latn form-of lang-pl nom|s-form-of        \" lang=\"pl\"><a rel=\"mw:WikiLink\" href=\"./Montreal#Polish\" class=\"mw-selflink-fragment\">Montreal</a></span><link rel=\"mw:PageProp/Category\" href=\"./Category:Polish_links_with_redundant_wikilinks#MONTREAL\"/></td></tr>\n<tr>\n<th title=\"dopełniacz (kogo? czego?)\" style=\"background:#eff7ff;\" scope=\"row\">genitive</th>\n<td><span class=\"Latn form-of lang-pl gen|s-form-of        \" lang=\"pl\"><a rel=\"mw:WikiLink\" href=\"./Montrealu?action=edit&amp;redlink=1#Polish\" title=\"Montrealu\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealu\"]}}'>Montrealu</a></span><link rel=\"mw:PageProp/Category\" href=\"./Category:Polish_links_with_redundant_wikilinks#MONTREAL\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Polish_links_with_redundant_alt_parameters#MONTREAL\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Polish_links_with_manual_fragments#MONTREAL\"/></td></tr>\n<tr>\n<th title=\"celownik (komu? czemu?)\" style=\"background:#eff7ff;\" scope=\"row\">dative</th>\n<td><span class=\"Latn form-of lang-pl dat|s-form-of        \" lang=\"pl\"><a rel=\"mw:WikiLink\" href=\"./Montrealowi?action=edit&amp;redlink=1#Polish\" title=\"Montrealowi\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealowi\"]}}'>Montrealowi</a></span><link rel=\"mw:PageProp/Category\" href=\"./Category:Polish_links_with_redundant_wikilinks#MONTREAL\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Polish_links_with_redundant_alt_parameters#MONTREAL\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Polish_links_with_manual_fragments#MONTREAL\"/></td></tr>\n<tr>\n<th title=\"biernik (kogo? co?)\" style=\"background:#eff7ff;\" scope=\"row\">accusative</th>\n<td><span class=\"Latn form-of lang-pl acc|s-form-of        \" lang=\"pl\"><a rel=\"mw:WikiLink\" href=\"./Montreal#Polish\" class=\"mw-selflink-fragment\">Montreal</a></span><link rel=\"mw:PageProp/Category\" href=\"./Category:Polish_links_with_redundant_wikilinks#MONTREAL\"/></td></tr>\n<tr>\n<th title=\"narzędnik (kim? czym?)\" style=\"background:#eff7ff;\" scope=\"row\">instrumental</th>\n<td><span class=\"Latn form-of lang-pl ins|s-form-of        \" lang=\"pl\"><a rel=\"mw:WikiLink\" href=\"./Montrealem?action=edit&amp;redlink=1#Polish\" title=\"Montrealem\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealem\"]}}'>Montrealem</a></span><link rel=\"mw:PageProp/Category\" href=\"./Category:Polish_links_with_redundant_wikilinks#MONTREAL\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Polish_links_with_redundant_alt_parameters#MONTREAL\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Polish_links_with_manual_fragments#MONTREAL\"/></td></tr>\n<tr>\n<th title=\"miejscownik (o kim? o czym?)\" style=\"background:#eff7ff;\" scope=\"row\">locative</th>\n<td><span class=\"Latn form-of lang-pl loc|s-form-of        \" lang=\"pl\"><a rel=\"mw:WikiLink\" href=\"./Montrealu?action=edit&amp;redlink=1#Polish\" title=\"Montrealu\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealu\"]}}'>Montrealu</a></span><link rel=\"mw:PageProp/Category\" href=\"./Category:Polish_links_with_redundant_wikilinks#MONTREAL\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Polish_links_with_redundant_alt_parameters#MONTREAL\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Polish_links_with_manual_fragments#MONTREAL\"/></td></tr>\n<tr>\n<th title=\"wołacz (o!)\" style=\"background:#eff7ff;\" scope=\"row\">vocative</th>\n<td><span class=\"Latn form-of lang-pl voc|s-form-of        \" lang=\"pl\"><a rel=\"mw:WikiLink\" href=\"./Montrealu?action=edit&amp;redlink=1#Polish\" title=\"Montrealu\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealu\"]}}'>Montrealu</a></span><link rel=\"mw:PageProp/Category\" href=\"./Category:Polish_links_with_redundant_wikilinks#MONTREAL\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Polish_links_with_redundant_alt_parameters#MONTREAL\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Polish_links_with_manual_fragments#MONTREAL\"/></td></tr>\n</tbody></table></div></div><link rel=\"mw:PageProp/Category\" href=\"./Category:Polish_singularia_tantum\" about=\"#mwt192\" id=\"mwAR8\"/>\n\n</section><section data-mw-section-id=\"30\" id=\"mwASA\"><h4 id=\"Derived_terms_2\">Derived terms</h4>\n<div class=\"term-list-header\" about=\"#mwt193\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"col-auto\",\"href\":\"./Template:col-auto\"},\"params\":{\"1\":{\"wt\":\"pl\"},\"title\":{\"wt\":\"adjective\"},\"2\":{\"wt\":\"montrealski\"}},\"i\":0}}]}' id=\"mwASE\">adjective</div><div class=\"list-switcher\" data-toggle-category=\"derived terms\" about=\"#mwt193\" id=\"mwASI\"><div class=\"derivedterms term-list ul-column-count\" data-column-count=\"1\" style=\"background-color:#F8F8FF\"><ul><li><span class=\"Latn\" lang=\"pl\"><a rel=\"mw:WikiLink\" href=\"./montrealski?action=edit&amp;redlink=1#Polish\" title=\"montrealski\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"montrealski\"]}}'>montrealski</a></span></li></ul></div><div class=\"list-switcher-element\" data-showtext=\" show more ▼ \" data-hidetext=\" show less ▲ \" style=\"display:none\"> </div></div>\n\n</section></section><section data-mw-section-id=\"31\" id=\"mwASM\"><h3 id=\"Further_reading\">Further reading</h3>\n<ul id=\"mwASQ\"><li id=\"mwASU\"><a rel=\"mw:ExtLink nofollow\" href=\"https://sjp.pwn.pl/szukaj/Montreal.html\" about=\"#mwt194\" typeof=\"mw:Transclusion\" class=\"external text\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"R:pl:PWN\",\"href\":\"./Template:R:pl:PWN\"},\"params\":{},\"i\":0}}]}' id=\"mwASY\">Montreal</a><span about=\"#mwt194\"> in Polish dictionaries at PWN</span></li></ul>\n\n</section></section><section data-mw-section-id=\"32\" id=\"mwASc\"><h2 id=\"Portuguese\">Portuguese</h2>\n\n<section data-mw-section-id=\"33\" id=\"mwASg\"><h3 id=\"Alternative_forms_2\">Alternative forms</h3>\n<ul id=\"mwASk\"><li id=\"mwASo\"><span class=\"Latn\" lang=\"pt\" about=\"#mwt195\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"alt\",\"href\":\"./Template:alt\"},\"params\":{\"1\":{\"wt\":\"pt\"},\"2\":{\"wt\":\"Montréal\"}},\"i\":0}}]}' id=\"mwASs\"><a rel=\"mw:WikiLink\" href=\"./Montréal#Portuguese\" title=\"Montréal\">Montréal</a></span></li></ul>\n\n</section><section data-mw-section-id=\"34\" id=\"mwASw\"><h3 id=\"Etymology_4\">Etymology</h3>\n<p id=\"mwAS0\">Borrowed from <span class=\"etyl\" about=\"#mwt196\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"bor\",\"href\":\"./Template:bor\"},\"params\":{\"1\":{\"wt\":\"pt\"},\"2\":{\"wt\":\"fr\"},\"3\":{\"wt\":\"Montréal\"}},\"i\":0}}]}' id=\"mwAS4\"><a rel=\"mw:WikiLink/Interwiki\" href=\"https://en.wikipedia.org/wiki/French%20language\" title=\"w:French language\" class=\"extiw\">French</a><link rel=\"mw:PageProp/Category\" href=\"./Category:Portuguese_terms_borrowed_from_French#MONTREAL\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Portuguese_terms_derived_from_French#MONTREAL\"/></span><span about=\"#mwt196\"> </span><i class=\"Latn mention\" lang=\"fr\" about=\"#mwt196\" id=\"mwAS8\"><a rel=\"mw:WikiLink\" href=\"./Montréal#French\" title=\"Montréal\">Montréal</a></i>.</p>\n\n</section><section data-mw-section-id=\"35\" id=\"mwATA\"><h3 id=\"Pronunciation_5\">Pronunciation</h3>\n<div about=\"#mwt197\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"pt-IPA\",\"href\":\"./Template:pt-IPA\"},\"params\":{},\"i\":0}}]}' id=\"mwATE\">\n<ul><li><span class=\"ib-brac qualifier-brac\">(</span><span class=\"ib-content qualifier-content\"><a rel=\"mw:WikiLink/Interwiki\" href=\"https://en.wikipedia.org/wiki/Brazilian_Portuguese\" title=\"w:Brazilian Portuguese\" class=\"extiw\">Brazil</a></span><span class=\"ib-brac qualifier-brac\">)</span> <a rel=\"mw:WikiLink\" href=\"./Wiktionary:International_Phonetic_Alphabet\" title=\"Wiktionary:International Phonetic Alphabet\">IPA</a><sup>(<a rel=\"mw:WikiLink\" href=\"./Appendix:Portuguese_pronunciation\" title=\"Appendix:Portuguese pronunciation\">key</a>)</sup>:<span typeof=\"mw:Entity\"> </span><span class=\"IPA\">/mõ.tɾeˈaw/</span> <span class=\"IPA\">[mõ.tɾeˈaʊ̯]</span><link rel=\"mw:PageProp/Category\" href=\"./Category:Portuguese_3-syllable_words\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Portuguese_terms_with_IPA_pronunciation#MONTREAL\"/></li></ul></div><span about=\"#mwt197\">\n</span><div class=\"vsSwitcher\" data-toggle-category=\"pronunciations\" style=\"width: 80em; max-width:100%;\" about=\"#mwt197\"><span class=\"vsToggleElement\" style=\"float: right;\"><span typeof=\"mw:Entity\"> </span></span>\n<div class=\"vsShow\" style=\"display:none\">\n<ul><li><span class=\"ib-brac qualifier-brac\">(</span><span class=\"ib-content qualifier-content\"><a rel=\"mw:WikiLink/Interwiki\" href=\"https://en.wikipedia.org/wiki/European_Portuguese\" title=\"w:European Portuguese\" class=\"extiw\">Portugal</a></span><span class=\"ib-brac qualifier-brac\">)</span> <a rel=\"mw:WikiLink\" href=\"./Wiktionary:International_Phonetic_Alphabet\" title=\"Wiktionary:International Phonetic Alphabet\">IPA</a><sup>(<a rel=\"mw:WikiLink\" href=\"./Appendix:Portuguese_pronunciation\" title=\"Appendix:Portuguese pronunciation\">key</a>)</sup>:<span typeof=\"mw:Entity\"> </span><span class=\"IPA\">/mõ.tɾiˈal/</span> <span class=\"IPA\">[mõ.tɾiˈaɫ]</span>, <span class=\"ib-brac qualifier-brac\">(</span><span class=\"ib-content qualifier-content\">faster pronunciation</span><span class=\"ib-brac qualifier-brac\">)</span> <span class=\"IPA\">/mõˈtɾjal/</span> <span class=\"IPA\">[mõˈtɾjaɫ]</span><link rel=\"mw:PageProp/Category\" href=\"./Category:Portuguese_3-syllable_words\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Portuguese_2-syllable_words\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Portuguese_terms_with_IPA_pronunciation#MONTREAL\"/></li></ul></div>\n<div class=\"vsHide\">\n<ul><li><span class=\"ib-brac qualifier-brac\">(</span><span class=\"ib-content qualifier-content\"><a rel=\"mw:WikiLink/Interwiki\" href=\"https://en.wikipedia.org/wiki/European_Portuguese\" title=\"w:European Portuguese\" class=\"extiw\">Portugal</a></span><span class=\"ib-brac qualifier-brac\">)</span> <a rel=\"mw:WikiLink\" href=\"./Wiktionary:International_Phonetic_Alphabet\" title=\"Wiktionary:International Phonetic Alphabet\">IPA</a><sup>(<a rel=\"mw:WikiLink\" href=\"./Appendix:Portuguese_pronunciation\" title=\"Appendix:Portuguese pronunciation\">key</a>)</sup>:<span typeof=\"mw:Entity\"> </span><span class=\"IPA\">/mõ.tɾiˈal/</span> <span class=\"IPA\">[mõ.tɾiˈaɫ]</span>, <span class=\"ib-brac qualifier-brac\">(</span><span class=\"ib-content qualifier-content\">faster pronunciation</span><span class=\"ib-brac qualifier-brac\">)</span> <span class=\"IPA\">/mõˈtɾjal/</span> <span class=\"IPA\">[mõˈtɾjaɫ]</span><link rel=\"mw:PageProp/Category\" href=\"./Category:Portuguese_3-syllable_words\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Portuguese_2-syllable_words\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Portuguese_terms_with_IPA_pronunciation#MONTREAL\"/>\n<ul><li><span class=\"ib-brac qualifier-brac\">(</span><span class=\"ib-content qualifier-content\">Southern Portugal</span><span class=\"ib-brac qualifier-brac\">)</span> <a rel=\"mw:WikiLink\" href=\"./Wiktionary:International_Phonetic_Alphabet\" title=\"Wiktionary:International Phonetic Alphabet\">IPA</a><sup>(<a rel=\"mw:WikiLink\" href=\"./Appendix:Portuguese_pronunciation\" title=\"Appendix:Portuguese pronunciation\">key</a>)</sup>:<span typeof=\"mw:Entity\"> </span><span class=\"IPA\">/mõ.tɾiˈa.li/</span>, <span class=\"ib-brac qualifier-brac\">(</span><span class=\"ib-content qualifier-content\">faster pronunciation</span><span class=\"ib-brac qualifier-brac\">)</span> <span class=\"IPA\">/mõˈtɾja.li/</span><link rel=\"mw:PageProp/Category\" href=\"./Category:Portuguese_4-syllable_words\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Portuguese_3-syllable_words\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Portuguese_terms_with_IPA_pronunciation#MONTREAL\"/></li></ul></li></ul></div></div><span about=\"#mwt197\">\n</span><p about=\"#mwt197\"><span></span></p>\n\n</section><section data-mw-section-id=\"36\" id=\"mwATI\"><h3 id=\"Proper_noun_8\">Proper noun</h3>\n<p id=\"mwATM\"><span class=\"headword-line\" about=\"#mwt198\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"pt-proper noun\",\"href\":\"./Template:pt-proper_noun\"},\"params\":{\"1\":{\"wt\":\"f\"}},\"i\":0}}]}' id=\"mwATQ\"><strong class=\"Latn headword\" lang=\"pt\">Montreal</strong><span typeof=\"mw:Entity\"> </span><span class=\"gender\"><abbr title=\"feminine gender\">f</abbr></span></span><link rel=\"mw:PageProp/Category\" href=\"./Category:Portuguese_lemmas#MONTREAL\" about=\"#mwt198\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Portuguese_proper_nouns#MONTREAL\" about=\"#mwt198\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Portuguese_feminine_nouns#MONTREAL\" about=\"#mwt198\" id=\"mwATU\"/></p>\n\n<ol id=\"mwATY\"><li id=\"mwATc\"><span class=\"Latn\" lang=\"en\" about=\"#mwt199\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"place\",\"href\":\"./Template:place\"},\"params\":{\"1\":{\"wt\":\"pt\"},\"2\":{\"wt\":\"largest city\"},\"3\":{\"wt\":\"p/Quebec\"},\"4\":{\"wt\":\"c/Canada\"},\"t1\":{\"wt\":\"Montreal\"}},\"i\":0}}]}' id=\"mwATg\"><a rel=\"mw:WikiLink\" href=\"./Montreal#English\" class=\"mw-selflink-fragment\">Montreal</a></span><span about=\"#mwt199\"> (the largest city in </span><span class=\"Latn\" lang=\"en\" about=\"#mwt199\"><a rel=\"mw:WikiLink\" href=\"./Quebec#English\" title=\"Quebec\">Quebec</a></span><span about=\"#mwt199\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt199\"><a rel=\"mw:WikiLink\" href=\"./Canada#English\" title=\"Canada\">Canada</a></span><span about=\"#mwt199\">)</span><link rel=\"mw:PageProp/Category\" href=\"./Category:pt:Montreal#MONTREAL\" about=\"#mwt199\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:pt:Montreal#MONTREAL\" about=\"#mwt199\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:pt:Cities_in_Quebec#MONTREAL\" about=\"#mwt199\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:pt:Cities_in_Canada#MONTREAL\" about=\"#mwt199\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:pt:Places_in_Quebec#MONTREAL\" about=\"#mwt199\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:pt:Places_in_Canada#MONTREAL\" about=\"#mwt199\" id=\"mwATk\"/></li></ol>\n\n</section></section><section data-mw-section-id=\"37\" id=\"mwATo\"><h2 id=\"Slovak\">Slovak</h2>\n\n<section data-mw-section-id=\"38\" id=\"mwATs\"><h3 id=\"Pronunciation_6\">Pronunciation</h3>\n<ul id=\"mwATw\"><li id=\"mwAT0\"><a rel=\"mw:WikiLink\" href=\"./Wiktionary:International_Phonetic_Alphabet\" title=\"Wiktionary:International Phonetic Alphabet\" about=\"#mwt200\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"sk-IPA\",\"href\":\"./Template:sk-IPA\"},\"params\":{},\"i\":0}}]}' id=\"mwAT4\">IPA</a><sup about=\"#mwt200\">(<a rel=\"mw:WikiLink/Interwiki\" href=\"https://en.wikipedia.org/wiki/Slovak%20phonology\" title=\"wikipedia:Slovak phonology\" class=\"extiw\">key</a>)</sup><span about=\"#mwt200\">:</span><span typeof=\"mw:Entity\" about=\"#mwt200\"> </span><span class=\"IPA\" about=\"#mwt200\">[ˈmɔntreal]</span><link rel=\"mw:PageProp/Category\" href=\"./Category:Slovak_terms_with_IPA_pronunciation#MONTREAL\" about=\"#mwt200\" id=\"mwAT8\"/></li></ul>\n\n</section><section data-mw-section-id=\"39\" id=\"mwAUA\"><h3 id=\"Proper_noun_9\">Proper noun</h3>\n<p id=\"mwAUE\"><span class=\"headword-line\" about=\"#mwt201\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"sk-proper noun\",\"href\":\"./Template:sk-proper_noun\"},\"params\":{\"1\":{\"wt\":\"m-in\"},\"gen\":{\"wt\":\"Montrealu\"},\"decl\":{\"wt\":\"dub\"}},\"i\":0}}]}' id=\"mwAUI\"><strong class=\"Latn headword\" lang=\"sk\">Montreal</strong><span typeof=\"mw:Entity\"> </span><span class=\"gender\"><abbr title=\"masculine gender\">m</abbr><span typeof=\"mw:Entity\"> </span><abbr title=\"inanimate\">inan</abbr></span> (<i>genitive singular</i> <b class=\"Latn\" lang=\"sk\"><a rel=\"mw:WikiLink\" href=\"./Montrealu?action=edit&amp;redlink=1#Slovak\" title=\"Montrealu\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealu\"]}}'>Montrealu</a></b>, <i>declension pattern of</i> <b class=\"Latn\" lang=\"sk\"><a rel=\"mw:WikiLink\" href=\"./dub#Slovak\" title=\"dub\">dub</a></b>)</span><link rel=\"mw:PageProp/Category\" href=\"./Category:Slovak_lemmas#MONTREAL\" about=\"#mwt201\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Slovak_proper_nouns#MONTREAL\" about=\"#mwt201\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Slovak_masculine_nouns#MONTREAL\" about=\"#mwt201\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Slovak_inanimate_nouns#MONTREAL\" about=\"#mwt201\" id=\"mwAUM\"/></p>\n\n<ol id=\"mwAUQ\"><li id=\"mwAUU\"><span class=\"Latn\" lang=\"en\" about=\"#mwt202\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"place\",\"href\":\"./Template:place\"},\"params\":{\"1\":{\"wt\":\"sk\"},\"2\":{\"wt\":\"largest city\"},\"3\":{\"wt\":\"p/Quebec\"},\"4\":{\"wt\":\"c/Canada\"},\"t\":{\"wt\":\"Montreal\"}},\"i\":0}}]}' id=\"mwAUY\"><a rel=\"mw:WikiLink\" href=\"./Montreal#English\" class=\"mw-selflink-fragment\">Montreal</a></span><span about=\"#mwt202\"> (the largest city in </span><span class=\"Latn\" lang=\"en\" about=\"#mwt202\"><a rel=\"mw:WikiLink\" href=\"./Quebec#English\" title=\"Quebec\">Quebec</a></span><span about=\"#mwt202\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt202\"><a rel=\"mw:WikiLink\" href=\"./Canada#English\" title=\"Canada\">Canada</a></span><span about=\"#mwt202\">)</span><link rel=\"mw:PageProp/Category\" href=\"./Category:sk:Montreal#MONTREAL\" about=\"#mwt202\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:sk:Montreal#MONTREAL\" about=\"#mwt202\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:sk:Cities_in_Quebec#MONTREAL\" about=\"#mwt202\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:sk:Cities_in_Canada#MONTREAL\" about=\"#mwt202\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:sk:Places_in_Quebec#MONTREAL\" about=\"#mwt202\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:sk:Places_in_Canada#MONTREAL\" about=\"#mwt202\" id=\"mwAUc\"/></li></ol>\n\n</section><section data-mw-section-id=\"40\" id=\"mwAUg\"><h3 id=\"References_2\">References</h3>\n<ul id=\"mwAUk\"><li id=\"mwAUo\"><span about=\"#mwt203\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"R:sk:SDK\",\"href\":\"./Template:R:sk:SDK\"},\"params\":{},\"i\":0}}]}' id=\"mwAUs\">“</span><a rel=\"mw:ExtLink nofollow\" href=\"https://slovnik.juls.savba.sk/?w=Montreal\" about=\"#mwt203\" class=\"external text\">Montreal</a><span about=\"#mwt203\">”, in</span><span typeof=\"mw:Entity\" about=\"#mwt203\"> </span><cite about=\"#mwt203\">Slovníkový portál Jazykovedného ústavu Ľ. Štúra SAV</cite><span typeof=\"mw:Entity\" about=\"#mwt203\"> </span><span typeof=\"mw:Entity\" about=\"#mwt203\">[</span><cite about=\"#mwt203\">Dictionary portal of the Ľ. Štúr Institute of Linguistics, Slovak Academy of Science</cite><span typeof=\"mw:Entity\" about=\"#mwt203\">]</span><span typeof=\"mw:Entity\" about=\"#mwt203\"> </span><span about=\"#mwt203\">(in Slovak), </span><a rel=\"mw:ExtLink nofollow\" href=\"https://slovnik.juls.savba.sk\" about=\"#mwt203\" class=\"external free\">https://slovnik.juls.savba.sk</a><span about=\"#mwt203\">, 2024</span></li></ul>\n\n</section></section><section data-mw-section-id=\"41\" id=\"mwAUw\"><h2 id=\"Spanish\">Spanish</h2>\n\n<section data-mw-section-id=\"42\" id=\"mwAU0\"><h3 id=\"Alternative_forms_3\">Alternative forms</h3>\n<ul id=\"mwAU4\"><li id=\"mwAU8\"><span class=\"Latn\" lang=\"es\" about=\"#mwt204\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"alt\",\"href\":\"./Template:alt\"},\"params\":{\"1\":{\"wt\":\"es\"},\"2\":{\"wt\":\"Montréal\"}},\"i\":0}}]}' id=\"mwAVA\"><a rel=\"mw:WikiLink\" href=\"./Montréal#Spanish\" title=\"Montréal\">Montréal</a></span></li></ul>\n\n</section><section data-mw-section-id=\"43\" id=\"mwAVE\"><h3 id=\"Pronunciation_7\">Pronunciation</h3>\n<ul about=\"#mwt205\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"es-pr\",\"href\":\"./Template:es-pr\"},\"params\":{},\"i\":0}}]}' id=\"mwAVI\"><li><a rel=\"mw:WikiLink\" href=\"./Wiktionary:International_Phonetic_Alphabet\" title=\"Wiktionary:International Phonetic Alphabet\">IPA</a><sup>(<a rel=\"mw:WikiLink\" href=\"./Appendix:Spanish_pronunciation\" title=\"Appendix:Spanish pronunciation\">key</a>)</sup>:<span typeof=\"mw:Entity\"> </span><span class=\"IPA\">/montɾeˈal/</span> <span class=\"IPA\">[mõn̪.t̪ɾeˈal]</span><link rel=\"mw:PageProp/Category\" href=\"./Category:Spanish_3-syllable_words\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Spanish_terms_with_IPA_pronunciation#MONTREAL\"/></li>\n<li>Rhymes: <a rel=\"mw:WikiLink\" href=\"./Rhymes:Spanish/al\" title=\"Rhymes:Spanish/al\"><span class=\"IPA\">-al</span></a><link rel=\"mw:PageProp/Category\" href=\"./Category:Rhymes:Spanish/al#MONTREAL\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Rhymes:Spanish/al/3_syllables#MONTREAL\"/></li>\n<li>Syllabification: <span class=\"Latn\" lang=\"es\">Mon‧tre‧al</span></li></ul>\n\n</section><section data-mw-section-id=\"44\" id=\"mwAVM\"><h3 id=\"Proper_noun_10\">Proper noun</h3>\n<p id=\"mwAVQ\"><span class=\"headword-line\" about=\"#mwt206\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"es-proper noun\",\"href\":\"./Template:es-proper_noun\"},\"params\":{\"1\":{\"wt\":\"f\"}},\"i\":0}}]}' id=\"mwAVU\"><strong class=\"Latn headword\" lang=\"es\">Montreal</strong><span typeof=\"mw:Entity\"> </span><span class=\"gender\"><abbr title=\"feminine gender\">f</abbr></span></span><link rel=\"mw:PageProp/Category\" href=\"./Category:Spanish_lemmas#MONTREAL\" about=\"#mwt206\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Spanish_proper_nouns#MONTREAL\" about=\"#mwt206\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Spanish_feminine_nouns#MONTREAL\" about=\"#mwt206\" id=\"mwAVY\"/></p>\n\n<ol id=\"mwAVc\"><li id=\"mwAVg\"><span class=\"Latn\" lang=\"en\" about=\"#mwt207\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"place\",\"href\":\"./Template:place\"},\"params\":{\"1\":{\"wt\":\"es\"},\"2\":{\"wt\":\"largest city\"},\"3\":{\"wt\":\"p/Quebec\"},\"4\":{\"wt\":\"c/Canada\"},\"t1\":{\"wt\":\"Montreal\"}},\"i\":0}}]}' id=\"mwAVk\"><a rel=\"mw:WikiLink\" href=\"./Montreal#English\" class=\"mw-selflink-fragment\">Montreal</a></span><span about=\"#mwt207\"> (the largest city in </span><span class=\"Latn\" lang=\"en\" about=\"#mwt207\"><a rel=\"mw:WikiLink\" href=\"./Quebec#English\" title=\"Quebec\">Quebec</a></span><span about=\"#mwt207\">, </span><span class=\"Latn\" lang=\"en\" about=\"#mwt207\"><a rel=\"mw:WikiLink\" href=\"./Canada#English\" title=\"Canada\">Canada</a></span><span about=\"#mwt207\">)</span><link rel=\"mw:PageProp/Category\" href=\"./Category:es:Montreal#MONTREAL\" about=\"#mwt207\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:es:Montreal#MONTREAL\" about=\"#mwt207\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:es:Cities_in_Quebec#MONTREAL\" about=\"#mwt207\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:es:Cities_in_Canada#MONTREAL\" about=\"#mwt207\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:es:Places_in_Quebec#MONTREAL\" about=\"#mwt207\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:es:Places_in_Canada#MONTREAL\" about=\"#mwt207\" id=\"mwAVo\"/></li></ol>\n\n</section></section><section data-mw-section-id=\"45\" id=\"mwAVs\"><h2 id=\"Turkish\">Turkish</h2>\n\n<section data-mw-section-id=\"46\" id=\"mwAVw\"><h3 id=\"Etymology_5\">Etymology</h3>\n<p id=\"mwAV0\">From <span class=\"etyl\" about=\"#mwt208\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"der\",\"href\":\"./Template:der\"},\"params\":{\"1\":{\"wt\":\"tr\"},\"2\":{\"wt\":\"fr\"},\"3\":{\"wt\":\"Montréal\"}},\"i\":0}}]}' id=\"mwAV4\"><a rel=\"mw:WikiLink/Interwiki\" href=\"https://en.wikipedia.org/wiki/French%20language\" title=\"w:French language\" class=\"extiw\">French</a><link rel=\"mw:PageProp/Category\" href=\"./Category:Turkish_terms_derived_from_French#MONTREAL\"/></span><span about=\"#mwt208\"> </span><i class=\"Latn mention\" lang=\"fr\" about=\"#mwt208\" id=\"mwAV8\"><a rel=\"mw:WikiLink\" href=\"./Montréal#French\" title=\"Montréal\">Montréal</a></i></p>\n\n</section><section data-mw-section-id=\"47\" id=\"mwAWA\"><h3 id=\"Pronunciation_8\">Pronunciation</h3>\n<ul id=\"mwAWE\"><li id=\"mwAWI\"><a rel=\"mw:WikiLink\" href=\"./Wiktionary:International_Phonetic_Alphabet\" title=\"Wiktionary:International Phonetic Alphabet\" about=\"#mwt209\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"IPA\",\"href\":\"./Template:IPA\"},\"params\":{\"1\":{\"wt\":\"tr\"},\"2\":{\"wt\":\"/mont.ɾeˈɑl/\"}},\"i\":0}}]}' id=\"mwAWM\">IPA</a><sup about=\"#mwt209\">(<a rel=\"mw:WikiLink\" href=\"./Appendix:Turkish_pronunciation\" title=\"Appendix:Turkish pronunciation\">key</a>)</sup><span about=\"#mwt209\">:</span><span typeof=\"mw:Entity\" about=\"#mwt209\"> </span><span class=\"IPA\" about=\"#mwt209\">/mont.ɾeˈɑl/</span><link rel=\"mw:PageProp/Category\" href=\"./Category:Turkish_terms_with_IPA_pronunciation#MONTREAL\" about=\"#mwt209\" id=\"mwAWQ\"/></li></ul>\n\n</section><section data-mw-section-id=\"48\" id=\"mwAWU\"><h3 id=\"Noun\">Noun</h3>\n<p id=\"mwAWY\"><span class=\"headword-line\" about=\"#mwt210\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"tr-noun\",\"href\":\"./Template:tr-noun\"},\"params\":{\"1\":{\"wt\":\"i\"},\"2\":{\"wt\":\"ler\"}},\"i\":0}}]}' id=\"mwAWc\"><strong class=\"Latn headword\" lang=\"tr\">Montreal</strong> (<i>definite accusative</i> <b class=\"Latn form-of lang-tr def|acc|s-form-of        \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreali?action=edit&amp;redlink=1#Turkish\" title=\"Montreali\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreali\"]}}'>Montreali</a></b>, <i>plural</i> <b class=\"Latn form-of lang-tr nom|p-form-of        \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealler?action=edit&amp;redlink=1#Turkish\" title=\"Montrealler\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealler\"]}}'>Montrealler</a></b>)</span><link rel=\"mw:PageProp/Category\" href=\"./Category:Turkish_lemmas#MONTREAL\" about=\"#mwt210\"/><link rel=\"mw:PageProp/Category\" href=\"./Category:Turkish_nouns#MONTREAL\" about=\"#mwt210\" id=\"mwAWg\"/></p>\n\n<ol id=\"mwAWk\"><li id=\"mwAWo\"><a rel=\"mw:WikiLink\" href=\"./Montreal\" class=\"mw-selflink selflink\" id=\"mwAWs\">Montreal</a></li></ol>\n\n<section data-mw-section-id=\"49\" id=\"mwAWw\"><h4 id=\"Declension_2\">Declension</h4>\n<table class=\"inflection-table vsSwitcher\" data-toggle-category=\"inflection\" style=\"text-align: left; background: #F9F9F9; border: 1px solid #AAAAAA;\" about=\"#mwt211\" typeof=\"mw:Transclusion\" data-mw='{\"parts\":[{\"template\":{\"target\":{\"wt\":\"tr-infl-noun-c\",\"href\":\"./Template:tr-infl-noun-c\"},\"params\":{\"1\":{\"wt\":\"e\"},\"poss\":{\"wt\":\"1\"}},\"i\":0}}]}' id=\"mwAW0\">\n<tbody><tr style=\"background: #DEDEDE; text-align: left;\">\n<th class=\"vsToggleElement\" colspan=\"3\">Inflection</th></tr>\n<tr class=\"vsShow\">\n<th style=\"background: #EFEFEF; width: 12em;\">Nominative</th>\n<td colspan=\"2\" style=\"width: 10em;\"><span class=\"Latn form-of lang-tr nom|s-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreal#Turkish\" class=\"mw-selflink-fragment\">Montreal</a></span></td></tr>\n<tr class=\"vsShow\">\n<th style=\"background: #EFEFEF;\">Definite accusative</th>\n<td colspan=\"2\"><span class=\"Latn form-of lang-tr def|acc|s-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreali?action=edit&amp;redlink=1#Turkish\" title=\"Montreali\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreali\"]}}'>Montreali</a></span></td></tr>\n<tr class=\"vsHide\" style=\"background: #DEDEDE;\">\n<th style=\"width: 12em;\"></th>\n<th style=\"width: 10em;\">Singular</th>\n<th style=\"width: 10em;\">Plural</th></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">Nominative</th>\n<td><span class=\"Latn form-of lang-tr nom|s-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreal#Turkish\" class=\"mw-selflink-fragment\">Montreal</a></span></td>\n<td><span class=\"Latn form-of lang-tr nom|p-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealler?action=edit&amp;redlink=1#Turkish\" title=\"Montrealler\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealler\"]}}'>Montrealler</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">Definite accusative</th>\n<td><span class=\"Latn form-of lang-tr def|acc|s-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreali?action=edit&amp;redlink=1#Turkish\" title=\"Montreali\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreali\"]}}'>Montreali</a></span></td>\n<td><span class=\"Latn form-of lang-tr def|acc|p-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealleri?action=edit&amp;redlink=1#Turkish\" title=\"Montrealleri\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealleri\"]}}'>Montrealleri</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">Dative</th>\n<td><span class=\"Latn form-of lang-tr dat|s-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreale?action=edit&amp;redlink=1#Turkish\" title=\"Montreale\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreale\"]}}'>Montreale</a></span></td>\n<td><span class=\"Latn form-of lang-tr dat|p-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallere?action=edit&amp;redlink=1#Turkish\" title=\"Montreallere\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallere\"]}}'>Montreallere</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">Locative</th>\n<td><span class=\"Latn form-of lang-tr loc|s-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealde?action=edit&amp;redlink=1#Turkish\" title=\"Montrealde\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealde\"]}}'>Montrealde</a></span></td>\n<td><span class=\"Latn form-of lang-tr loc|p-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerde?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerde\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerde\"]}}'>Montreallerde</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">Ablative</th>\n<td><span class=\"Latn form-of lang-tr abl|s-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealden?action=edit&amp;redlink=1#Turkish\" title=\"Montrealden\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealden\"]}}'>Montrealden</a></span></td>\n<td><span class=\"Latn form-of lang-tr abl|p-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerden?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerden\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerden\"]}}'>Montreallerden</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">Genitive</th>\n<td><span class=\"Latn form-of lang-tr gen|s-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealin?action=edit&amp;redlink=1#Turkish\" title=\"Montrealin\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealin\"]}}'>Montrealin</a></span></td>\n<td><span class=\"Latn form-of lang-tr gen|p-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerin?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerin\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerin\"]}}'>Montreallerin</a></span></td></tr>\n<tr class=\"vsHide\">\n<td colspan=\"3\" style=\"padding: 0;\">\n<table class=\"inflection-table vsSwitcher\" style=\"text-align: left; width: 100%;\">\n<tbody><tr style=\"background: #DEDEDE; text-align: left;\">\n<th class=\"vsToggleElement\" colspan=\"3\">Possessive forms</th></tr>\n<tr class=\"vsHide\" style=\"background: #DEDEDE;\">\n<th colspan=\"3\">Nominative</th></tr>\n<tr class=\"vsHide\" style=\"background: #DEDEDE;\">\n<th style=\"width: 12em;\"></th>\n<th style=\"width: 10em;\">Singular</th>\n<th style=\"width: 10em;\">Plural</th></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">1st singular</th>\n<td><span class=\"Latn form-of lang-tr 1|s|spos|poss-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealim?action=edit&amp;redlink=1#Turkish\" title=\"Montrealim\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealim\"]}}'>Montrealim</a></span></td>\n<td><span class=\"Latn form-of lang-tr 1|s|mpos|poss-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerim?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerim\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerim\"]}}'>Montreallerim</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">2nd singular</th>\n<td><span class=\"Latn form-of lang-tr 2|s|spos|poss-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealin?action=edit&amp;redlink=1#Turkish\" title=\"Montrealin\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealin\"]}}'>Montrealin</a></span></td>\n<td><span class=\"Latn form-of lang-tr 2|s|mpos|poss-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerin?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerin\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerin\"]}}'>Montreallerin</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">3rd singular</th>\n<td><span class=\"Latn form-of lang-tr 3|s|spos|poss-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreali?action=edit&amp;redlink=1#Turkish\" title=\"Montreali\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreali\"]}}'>Montreali</a></span></td>\n<td><span class=\"Latn form-of lang-tr 3|s|mpos|poss-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealleri?action=edit&amp;redlink=1#Turkish\" title=\"Montrealleri\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealleri\"]}}'>Montrealleri</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">1st plural</th>\n<td><span class=\"Latn form-of lang-tr 1|p|spos|poss-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealimiz?action=edit&amp;redlink=1#Turkish\" title=\"Montrealimiz\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealimiz\"]}}'>Montrealimiz</a></span></td>\n<td><span class=\"Latn form-of lang-tr 1|p|mpos|poss-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerimiz?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerimiz\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerimiz\"]}}'>Montreallerimiz</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">2nd plural</th>\n<td><span class=\"Latn form-of lang-tr 2|p|spos|poss-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealiniz?action=edit&amp;redlink=1#Turkish\" title=\"Montrealiniz\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealiniz\"]}}'>Montrealiniz</a></span></td>\n<td><span class=\"Latn form-of lang-tr 2|p|mpos|poss-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealleriniz?action=edit&amp;redlink=1#Turkish\" title=\"Montrealleriniz\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealleriniz\"]}}'>Montrealleriniz</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">3rd plural</th>\n<td><span class=\"Latn form-of lang-tr 3|p|spos|poss-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealleri?action=edit&amp;redlink=1#Turkish\" title=\"Montrealleri\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealleri\"]}}'>Montrealleri</a></span></td>\n<td><span class=\"Latn form-of lang-tr 3|p|mpos|poss-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealleri?action=edit&amp;redlink=1#Turkish\" title=\"Montrealleri\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealleri\"]}}'>Montrealleri</a></span></td></tr>\n<tr class=\"vsHide\" style=\"background: #DEDEDE;\">\n<th colspan=\"3\">Definite accusative</th></tr>\n<tr class=\"vsHide\" style=\"background: #DEDEDE;\">\n<th style=\"width: 12em;\"></th>\n<th style=\"width: 10em;\">Singular</th>\n<th style=\"width: 10em;\">Plural</th></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">1st singular</th>\n<td><span class=\"Latn form-of lang-tr 1|s|spos|poss|def|acc-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealimi?action=edit&amp;redlink=1#Turkish\" title=\"Montrealimi\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealimi\"]}}'>Montrealimi</a></span></td>\n<td><span class=\"Latn form-of lang-tr 1|s|mpos|poss|def|acc-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerimi?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerimi\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerimi\"]}}'>Montreallerimi</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">2nd singular</th>\n<td><span class=\"Latn form-of lang-tr 2|s|spos|poss|def|acc-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealini?action=edit&amp;redlink=1#Turkish\" title=\"Montrealini\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealini\"]}}'>Montrealini</a></span></td>\n<td><span class=\"Latn form-of lang-tr 2|s|mpos|poss|def|acc-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerini?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerini\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerini\"]}}'>Montreallerini</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">3rd singular</th>\n<td><span class=\"Latn form-of lang-tr 3|s|spos|poss|def|acc-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealini?action=edit&amp;redlink=1#Turkish\" title=\"Montrealini\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealini\"]}}'>Montrealini</a></span></td>\n<td><span class=\"Latn form-of lang-tr 3|s|mpos|poss|def|acc-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerini?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerini\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerini\"]}}'>Montreallerini</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">1st plural</th>\n<td><span class=\"Latn form-of lang-tr 1|p|spos|poss|def|acc-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealimizi?action=edit&amp;redlink=1#Turkish\" title=\"Montrealimizi\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealimizi\"]}}'>Montrealimizi</a></span></td>\n<td><span class=\"Latn form-of lang-tr 1|p|mpos|poss|def|acc-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerimizi?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerimizi\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerimizi\"]}}'>Montreallerimizi</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">2nd plural</th>\n<td><span class=\"Latn form-of lang-tr 2|p|spos|poss|def|acc-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealinizi?action=edit&amp;redlink=1#Turkish\" title=\"Montrealinizi\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealinizi\"]}}'>Montrealinizi</a></span></td>\n<td><span class=\"Latn form-of lang-tr 2|p|mpos|poss|def|acc-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerinizi?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerinizi\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerinizi\"]}}'>Montreallerinizi</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">3rd plural</th>\n<td><span class=\"Latn form-of lang-tr 3|p|spos|poss|def|acc-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerini?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerini\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerini\"]}}'>Montreallerini</a></span></td>\n<td><span class=\"Latn form-of lang-tr 3|p|mpos|poss|def|acc-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerini?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerini\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerini\"]}}'>Montreallerini</a></span></td></tr>\n<tr class=\"vsHide\" style=\"background: #DEDEDE;\">\n<th colspan=\"3\">Dative</th></tr>\n<tr class=\"vsHide\" style=\"background: #DEDEDE;\">\n<th style=\"width: 12em;\"></th>\n<th style=\"width: 10em;\">Singular</th>\n<th style=\"width: 10em;\">Plural</th></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">1st singular</th>\n<td><span class=\"Latn form-of lang-tr 1|s|spos|poss|dat-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealime?action=edit&amp;redlink=1#Turkish\" title=\"Montrealime\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealime\"]}}'>Montrealime</a></span></td>\n<td><span class=\"Latn form-of lang-tr 1|s|mpos|poss|dat-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerime?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerime\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerime\"]}}'>Montreallerime</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">2nd singular</th>\n<td><span class=\"Latn form-of lang-tr 2|s|spos|poss|dat-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealine?action=edit&amp;redlink=1#Turkish\" title=\"Montrealine\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealine\"]}}'>Montrealine</a></span></td>\n<td><span class=\"Latn form-of lang-tr 2|s|mpos|poss|dat-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerine?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerine\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerine\"]}}'>Montreallerine</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">3rd singular</th>\n<td><span class=\"Latn form-of lang-tr 3|s|spos|poss|dat-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealine?action=edit&amp;redlink=1#Turkish\" title=\"Montrealine\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealine\"]}}'>Montrealine</a></span></td>\n<td><span class=\"Latn form-of lang-tr 3|s|mpos|poss|dat-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerine?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerine\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerine\"]}}'>Montreallerine</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">1st plural</th>\n<td><span class=\"Latn form-of lang-tr 1|p|spos|poss|dat-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealimize?action=edit&amp;redlink=1#Turkish\" title=\"Montrealimize\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealimize\"]}}'>Montrealimize</a></span></td>\n<td><span class=\"Latn form-of lang-tr 1|p|mpos|poss|dat-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerimize?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerimize\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerimize\"]}}'>Montreallerimize</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">2nd plural</th>\n<td><span class=\"Latn form-of lang-tr 2|p|spos|poss|dat-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealinize?action=edit&amp;redlink=1#Turkish\" title=\"Montrealinize\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealinize\"]}}'>Montrealinize</a></span></td>\n<td><span class=\"Latn form-of lang-tr 2|p|mpos|poss|dat-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerinize?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerinize\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerinize\"]}}'>Montreallerinize</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">3rd plural</th>\n<td><span class=\"Latn form-of lang-tr 3|p|spos|poss|dat-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerine?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerine\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerine\"]}}'>Montreallerine</a></span></td>\n<td><span class=\"Latn form-of lang-tr 3|p|mpos|poss|dat-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerine?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerine\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerine\"]}}'>Montreallerine</a></span></td></tr>\n<tr class=\"vsHide\" style=\"background: #DEDEDE;\">\n<th colspan=\"3\">Locative</th></tr>\n<tr class=\"vsHide\" style=\"background: #DEDEDE;\">\n<th style=\"width: 12em;\"></th>\n<th style=\"width: 10em;\">Singular</th>\n<th style=\"width: 10em;\">Plural</th></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">1st singular</th>\n<td><span class=\"Latn form-of lang-tr 1|s|spos|poss|loc-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealimde?action=edit&amp;redlink=1#Turkish\" title=\"Montrealimde\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealimde\"]}}'>Montrealimde</a></span></td>\n<td><span class=\"Latn form-of lang-tr 1|s|mpos|poss|loc-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerimde?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerimde\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerimde\"]}}'>Montreallerimde</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">2nd singular</th>\n<td><span class=\"Latn form-of lang-tr 2|s|spos|poss|loc-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealinde?action=edit&amp;redlink=1#Turkish\" title=\"Montrealinde\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealinde\"]}}'>Montrealinde</a></span></td>\n<td><span class=\"Latn form-of lang-tr 2|s|mpos|poss|loc-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerinde?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerinde\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerinde\"]}}'>Montreallerinde</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">3rd singular</th>\n<td><span class=\"Latn form-of lang-tr 3|s|spos|poss|loc-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealinde?action=edit&amp;redlink=1#Turkish\" title=\"Montrealinde\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealinde\"]}}'>Montrealinde</a></span></td>\n<td><span class=\"Latn form-of lang-tr 3|s|mpos|poss|loc-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerinde?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerinde\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerinde\"]}}'>Montreallerinde</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">1st plural</th>\n<td><span class=\"Latn form-of lang-tr 1|p|spos|poss|loc-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealimizde?action=edit&amp;redlink=1#Turkish\" title=\"Montrealimizde\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealimizde\"]}}'>Montrealimizde</a></span></td>\n<td><span class=\"Latn form-of lang-tr 1|p|mpos|poss|loc-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerimizde?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerimizde\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerimizde\"]}}'>Montreallerimizde</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">2nd plural</th>\n<td><span class=\"Latn form-of lang-tr 2|p|spos|poss|loc-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealinizde?action=edit&amp;redlink=1#Turkish\" title=\"Montrealinizde\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealinizde\"]}}'>Montrealinizde</a></span></td>\n<td><span class=\"Latn form-of lang-tr 2|p|mpos|poss|loc-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerinizde?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerinizde\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerinizde\"]}}'>Montreallerinizde</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">3rd plural</th>\n<td><span class=\"Latn form-of lang-tr 3|p|spos|poss|loc-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerinde?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerinde\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerinde\"]}}'>Montreallerinde</a></span></td>\n<td><span class=\"Latn form-of lang-tr 3|p|mpos|poss|loc-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerinde?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerinde\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerinde\"]}}'>Montreallerinde</a></span></td></tr>\n<tr class=\"vsHide\" style=\"background: #DEDEDE;\">\n<th colspan=\"3\">Ablative</th></tr>\n<tr class=\"vsHide\" style=\"background: #DEDEDE;\">\n<th style=\"width: 12em;\"></th>\n<th style=\"width: 10em;\">Singular</th>\n<th style=\"width: 10em;\">Plural</th></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">1st singular</th>\n<td><span class=\"Latn form-of lang-tr 1|s|spos|poss|abl-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealimden?action=edit&amp;redlink=1#Turkish\" title=\"Montrealimden\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealimden\"]}}'>Montrealimden</a></span></td>\n<td><span class=\"Latn form-of lang-tr 1|s|mpos|poss|abl-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerimden?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerimden\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerimden\"]}}'>Montreallerimden</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">2nd singular</th>\n<td><span class=\"Latn form-of lang-tr 2|s|spos|poss|abl-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealinden?action=edit&amp;redlink=1#Turkish\" title=\"Montrealinden\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealinden\"]}}'>Montrealinden</a></span></td>\n<td><span class=\"Latn form-of lang-tr 2|s|mpos|poss|abl-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerinden?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerinden\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerinden\"]}}'>Montreallerinden</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">3rd singular</th>\n<td><span class=\"Latn form-of lang-tr 3|s|spos|poss|abl-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealinden?action=edit&amp;redlink=1#Turkish\" title=\"Montrealinden\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealinden\"]}}'>Montrealinden</a></span></td>\n<td><span class=\"Latn form-of lang-tr 3|s|mpos|poss|abl-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerinden?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerinden\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerinden\"]}}'>Montreallerinden</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">1st plural</th>\n<td><span class=\"Latn form-of lang-tr 1|p|spos|poss|abl-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealimizden?action=edit&amp;redlink=1#Turkish\" title=\"Montrealimizden\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealimizden\"]}}'>Montrealimizden</a></span></td>\n<td><span class=\"Latn form-of lang-tr 1|p|mpos|poss|abl-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerimizden?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerimizden\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerimizden\"]}}'>Montreallerimizden</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">2nd plural</th>\n<td><span class=\"Latn form-of lang-tr 2|p|spos|poss|abl-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealinizden?action=edit&amp;redlink=1#Turkish\" title=\"Montrealinizden\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealinizden\"]}}'>Montrealinizden</a></span></td>\n<td><span class=\"Latn form-of lang-tr 2|p|mpos|poss|abl-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerinizden?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerinizden\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerinizden\"]}}'>Montreallerinizden</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">3rd plural</th>\n<td><span class=\"Latn form-of lang-tr 3|p|spos|poss|abl-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerinden?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerinden\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerinden\"]}}'>Montreallerinden</a></span></td>\n<td><span class=\"Latn form-of lang-tr 3|p|mpos|poss|abl-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerinden?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerinden\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerinden\"]}}'>Montreallerinden</a></span></td></tr>\n<tr class=\"vsHide\" style=\"background: #DEDEDE;\">\n<th colspan=\"3\">Genitive</th></tr>\n<tr class=\"vsHide\" style=\"background: #DEDEDE;\">\n<th style=\"width: 12em;\"></th>\n<th style=\"width: 10em;\">Singular</th>\n<th style=\"width: 10em;\">Plural</th></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">1st singular</th>\n<td><span class=\"Latn form-of lang-tr 1|s|spos|poss|gen-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealimin?action=edit&amp;redlink=1#Turkish\" title=\"Montrealimin\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealimin\"]}}'>Montrealimin</a></span></td>\n<td><span class=\"Latn form-of lang-tr 1|s|mpos|poss|gen-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerimin?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerimin\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerimin\"]}}'>Montreallerimin</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">2nd singular</th>\n<td><span class=\"Latn form-of lang-tr 2|s|spos|poss|gen-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealinin?action=edit&amp;redlink=1#Turkish\" title=\"Montrealinin\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealinin\"]}}'>Montrealinin</a></span></td>\n<td><span class=\"Latn form-of lang-tr 2|s|mpos|poss|gen-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerinin?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerinin\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerinin\"]}}'>Montreallerinin</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">3rd singular</th>\n<td><span class=\"Latn form-of lang-tr 3|s|spos|poss|gen-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealinin?action=edit&amp;redlink=1#Turkish\" title=\"Montrealinin\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealinin\"]}}'>Montrealinin</a></span></td>\n<td><span class=\"Latn form-of lang-tr 3|s|mpos|poss|gen-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerinin?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerinin\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerinin\"]}}'>Montreallerinin</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">1st plural</th>\n<td><span class=\"Latn form-of lang-tr 1|p|spos|poss|gen-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealimizin?action=edit&amp;redlink=1#Turkish\" title=\"Montrealimizin\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealimizin\"]}}'>Montrealimizin</a></span></td>\n<td><span class=\"Latn form-of lang-tr 1|p|mpos|poss|gen-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerimizin?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerimizin\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerimizin\"]}}'>Montreallerimizin</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">2nd plural</th>\n<td><span class=\"Latn form-of lang-tr 2|p|spos|poss|gen-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montrealinizin?action=edit&amp;redlink=1#Turkish\" title=\"Montrealinizin\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montrealinizin\"]}}'>Montrealinizin</a></span></td>\n<td><span class=\"Latn form-of lang-tr 2|p|mpos|poss|gen-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerinizin?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerinizin\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerinizin\"]}}'>Montreallerinizin</a></span></td></tr>\n<tr class=\"vsHide\">\n<th style=\"background: #EFEFEF;\">3rd plural</th>\n<td><span class=\"Latn form-of lang-tr 3|p|spos|poss|gen-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerinin?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerinin\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerinin\"]}}'>Montreallerinin</a></span></td>\n<td><span class=\"Latn form-of lang-tr 3|p|mpos|poss|gen-form-of     origin-Montreal   \" lang=\"tr\"><a rel=\"mw:WikiLink\" href=\"./Montreallerinin?action=edit&amp;redlink=1#Turkish\" title=\"Montreallerinin\" class=\"new\" typeof=\"mw:LocalizedAttrs\" data-mw-i18n='{\"title\":{\"lang\":\"x-page\",\"key\":\"red-link-title\",\"params\":[\"Montreallerinin\"]}}'>Montreallerinin</a></span></td></tr>\n</tbody></table></td></tr>\n</tbody></table></section></section></section></body></html>",
            "wikitext": "{{also|Montréal|Mont-real}}\n==English==\n{{wikipedia}}\n{{commonscat}}\n[[Image:Montreal Twilight Panorama 2006.jpg|thumb|right|A panorama of Montreal.]]\n\n===Alternative forms===\n* {{qualifier|official}} {{alter|en|Montréal}}\n\n===Etymology===\nFrom {{bor|en|fr|Montréal}}; from {{m|fr|Mont Royal|gloss=[[Mount Royal]]}} {{gloss|{{compound|fr|mont|royal|nocat=1}}}}, after the nearby hill by Jacques Cartier. It is uncertain how {{m|fr|royal|Royal}} became ''-real''. A common explanation is that {{m|frm|real}} is the {{der|en|frm|-}} form of ''Royal'', but Cartier himself recorded ''le mont Royal'' when he named it. It may have perhaps been from the Italian G.B. Rasmusio's 1556 map translating the name to {{der|en|it|Monte Real}}.<ref>{{cite-web |website= Natural Resources Canada |title= Geographical Name - Island of Montreal |date= 22 August 2014 |url=http://www.nrcan.gc.ca/earth-sciences/geography/place-names/origins-geographical-names/9218 |accessdate=15 August 2015 |archiveurl=https://web.archive.org/web/20160803212245/https://www.nrcan.gc.ca/earth-sciences/geography/place-names/origins-geographical-names/9218 |archivedate=3 August 2016 |author= Jean Poirier |publicationdate= 1979 }}</ref>\n\n===Pronunciation===\n* {{a|UK}} {{IPA|en|/ˌmɒntɹiˈɔːl/}}\n* {{a|US}} {{IPA|en|/ˌmɑntɹiˈɔl/}}\n* {{a|CA}} {{IPA|en|/ˌmʌntɹiːˈɒl/|/ˌmɔ̃ɹeɪˈæl/}}\n** {{audio|en|Montreal-English-pronunciation.oga|Audio (Canada)}}\n* {{rhymes|en|ɔːl|s=3}}\n\n===Proper noun===\n{{en-proper noun}}\n\n# {{place|en|place|p/Quebec|c/Canada}}, at the confluence of the Ottawa and Saint-Lawrence Rivers\n## {{place|en|A river <<port>>, the <<largest city>> in <<p/Quebec>>, <<c/Canada>>}}.\n##: {{syn|en|City of Montreal}}\n## {{place|en|island|p/Quebec|c/Canada}}, on which the city is situated.\n##: {{syn|en|Island of Montreal|Montreal Island}}\n##: {{cot|en|514}}\n## {{ellipsis of|en|Montreal Archipelago}}; {{place|en|archipelago|on the|river/Saint Lawrence River|in|p/Quebec|c/Canada}}, which contains the island.\n##: {{syn|en|Hochelaga Archipelago}}\n## {{place|en|administrative region|p/Quebec|c/Canada}}, containing the island.\n##: {{syn|en|Montreal Region|Region of Montreal|Region 6|Region 06}}\n## {{place|en|metropolitan area|p/Quebec|c/Canada}}, containing the archipelago.\n##: {{hypo|en|Montreal Metropolitan Community|Montreal Urban Community|Agglomeration of Montreal|Greater Montreal Area|Greater Montreal|Metro Montreal}}\n## {{lb|en|historical}} {{place|en|former county|p/Quebec|c/Canada}}, containing the archipelago.\n##: {{syn|en|County of Montreal|Montreal County}}\n## {{place|en|district|p/Quebec|c/Canada}}, containing the county.\n##: {{syn|en|District of Montreal|Montreal District}}\n### {{lb|en|judicial}} {{place|en|judicial district|p/Quebec|c/Canada}}\n### {{lb|en|historical}} {{place|en|former administrative district|p/Quebec|c/Canada}}\n## {{lb|en|historical}} {{place|en|district|p/Quebec|cc/British North America}}\n## {{lb|en|historical}} {{place|en|district|p/Canada|cc/New France}}\n# {{place|en|river|dist/Algoma,Sudbury|p/Ontario|c/Canada}}.\n# {{place|en|river|dist/Timiskaming|p/Ontario|c/Canada}}.\n# {{place|en|river|cdiv/No. 18|p/Saskatchewan|c/Canada}}.\n# {{place|en|unincorporated community|co/Camden County|s/Missouri|c/USA}}.\n# {{place|en|small city|co/Iron County|s/Wisconsin|c/USA}}.\n# {{place|en|river|on the|pen:Suf/Keweenaw|s/Michigan|c/USA}}.\n# {{place|en|A <<river>> in northern <<s/Wisconsin>> and the <<pen/Upper Peninsula>> of <<s/Michigan>>, <<c/USA>>}}.\n\n====Synonyms====\n* {{l|en|Mtl.}}, {{l|en|Mtl}}, {{l|en|MTL}}; {{gloss|abbreviation}}\n\n====Derived terms====\n{{der-top}}\n* {{l|en|Montrealer}}\n* {{l|en|Montreal-style}}\n* {{l|en|Montreal bagel}}\n* {{l|en|Montreal smoked meat}}\n* {{l|en|Montreal staircase}}\n* {{l|en|Montreal North}}\n* {{l|en|Montreal nutmeg melon}}\n*{{l|en|Montreal Protocol}}\n* {{l|en|Montreal South}}\n* {{l|en|Montreal East}}\n* {{l|en|Montreal West}}\n* {{l|en|Old Montreal}}\n{{der-bottom}}\n\n====Translations====\n{{trans-top|a city in Quebec}}\n* Arabic: {{t|ar|مُونْتْرِيَال|m|tr=montreāl}}\n* Armenian: {{t|hy|Մոնրեալ}}\n* Belarusian: {{t|be|Манрэа́ль|m}}\n* Bulgarian: {{t|bg|Монреа́л|m}}\n* Catalan: {{t+|ca|Mont-real|m}}\n* Chinese:\n*: Cantonese: {{t|yue|滿地可|tr=mun5 dei6 ho2}}\n*: Mandarin: {{t+|cmn|蒙特利爾|tr=Méngtèlì'ěr}}, {{t|cmn|蒙特婁|tr=Méngtèlóu}}, {{t|cmn|滿地河|tr=Mǎndìhé}}, {{t+|cmn|蒙城|tr=Méngchéng}} {{qualifier|short form}}\n* Esperanto: {{t|eo|Montrealo}}\n* Farefare: {{t|gur|Mõrɛallɛ}}\n* Finnish: {{t+|fi|Montreal}}\n* French: {{t+|fr|Montréal|m}}\n*: Middle French: {{t|frm|Place Royale}}, {{t|frm|Ville Marie}}\n* Georgian: {{t|ka|მონრეალი}}\n* German: {{t+|de|Montreal|n}}, {{t+|de|Montréal|n}}\n* Greek: {{t+|el|Μόντρεαλ}}, {{t|el|Μονρεάλη}}\n* Gujarati: {{t|gu|મોન્ટ્રીયલ}}\n* Haitian Creole: {{t|ht|Monreyal}}\n* Hebrew: {{t|he|מונטריאול}}\n* Hindi: {{t|hi|मॉन्ट्रियल}}\n* Interlingua: {{t|ia|Montreal}}\n* Inuktitut: {{t|iu|ᒧᓐᕆᐊᓪ|sc=Cans}}, {{t|iu|ᒧᓐᑐᕆᐋᓪ|sc=Cans}}\n* Italian: {{t|it|Montreale|f}}\n* Japanese: {{t|ja|モントリオール|tr=Montoriōru}}, {{t|ja|モンレアル|tr=Monrearu}}\n* Khmer: {{t|km|ម៉ុងរ៉េអាល់}}\n* Korean: {{t+|ko|^몬트리올}}\n* Latin: {{t|la|Mons Rēgius|m}}, {{t|la|Mōns Rēgālis|m}}, {{t|la|Marianopolis|m}}\n* Latvian: {{t|lv|Monreala|f}}\n* Laurentian: {{t|lre|Tutonaguy}}\n* Lithuanian: {{t|lt|Monrealis}}\n* Macedonian: {{t|mk|Монтреа́л|m}}\n* Norman: {{t|nrf|Mountriyal}}\n* Norwegian:\n*: Bokmål: {{t|nb|Montréal}}, {{t|nb|Montreal}}\n*: Nynorsk: {{t|nn|Montréal}}, {{t|nn|Montreal}}\n* Ojibwe: {{t|oj|Mooniyaang}}\n* Ottawa: {{t|otw|Moon'yaang}}\n* Persian: {{t|fa|مونترال}}\n* Polish: {{t+|pl|Montreal|m-in}}\n* Portuguese: {{t|pt|Montreal|f}}, {{t|pt|Montréal|f}}\n* Romanian: {{t|ro|Montreal|n}}\n* Russian: {{t+|ru|Монреа́ль|m}}\n* Scottish Gaelic: {{t|gd|Beinn Rìoghail|f}}\n* Serbo-Croatian:\n*: Cyrillic: {{t|sh|Монтреал|m}}\n*: Roman: {{t+|sh|Montreal|m}}\n* Spanish: {{t+|es|Montreal|m}}\n* Thai: {{t|th|มอนทรีออล}}\n* Turkish: {{t|tr|Montreal}}\n* Ukrainian: {{t|uk|Монреа́ль|m}}\n* Vietnamese: {{t|vi|Mông-rê-an}}, {{t|vi|Montréal}}\n* Yiddish: {{t|yi|מאָנטרעאָל}}, {{t|yi|מאָנטרעאַל}}\n{{trans-bottom}}\n\n{{trans-top|an island in Quebec}}\n* French: {{t+|fr|Montréal}}\n*: Middle French: {{t|frm|Montreal}}\n* Laurentian: {{t|lre|Hochelaga}}\n{{trans-bottom}}\n\n===References===\n<references/>\n\n===Anagrams===\n* {{anagrams|en|a=aelmnort|loan term|morantel}}\n\n==Catalan==\n{{wikipedia|lang=ca}}\n\n===Pronunciation===\n* {{ca-IPA|[Mónt.,àl]}}<!-- per esAdir -->\n\n===Proper noun===\n{{ca-proper noun|m|head=[[mont|Mont]][[real]]}}\n\n# {{place|ca|largest city|p/Quebec|c/Canada|t1=Montreal}}\n# {{place|ca|island|p/Quebec|c/Canada|the site of the city of [[Montreal]]|t1=Montreal}}\n\n==German==\n{{wikipedia|lang=de}}\n\n===Pronunciation===\n* {{IPA|de|[mɔntʁeˈaːl]}}\n* {{audio|de|De-Montreal.ogg|Audio}}\n\n===Proper noun===\n{{de-proper noun|toponym}}\n\n# {{place|de|largest city|p/Quebec|c/Canada|t1=Montreal}}\n\n==Middle French==\n\n===Etymology===\n{{compound|frm|mont|real}}; The island was named after the mountain located on it, {{m|frm|mont Royal}}, which was named by Jacques Cartier.\n\n===Proper noun===\n{{head|frm|proper noun}}\n\n# {{place|frm|island|riv:Suf/Saint Lawrence|in modern|c/Canada}}, at its confluence with the [[Ottawa]] River.\n\n====Descendants====\n{{top2}}\n* {{desc|en|Montreal}}\n* {{desc|fr|Montréal}}\n{{bottom}}\n\n==Norwegian Bokmål==\n\n===Proper noun===\n{{head|nb|proper noun}}\n\n# {{alt form|nb|Montréal|nocap=1}}, {{place|nb|a=the|largest city|p/Quebec|c/Canada}}\n\n==Norwegian Nynorsk==\n\n===Proper noun===\n{{head|nn|proper noun}}\n\n# {{alt form|nn|Montréal|nocap=1}}, {{place|nn|a=the|largest city|p/Quebec|c/Canada}}\n\n==Polish==\n{{wp|lang=pl}}\n\n===Etymology===\n{{bor+|pl|en|Montreal}}, from {{der|pl|fr|Montréal}}, from {{m|fr|Mont Royal|gloss=[[Mount Royal]]}}, after the nearby mountain by Jacques Cartier.\n\n===Pronunciation===\n{{pl-p|a=Pl-Montreal.ogg}}\n\n===Proper noun===\n{{pl-proper noun|m-in}}\n\n# {{place|pl|largest city|p/Quebec|c/Canada|t1=Montreal}}\n\n====Declension====\n{{pl-decl-noun-m-in|tantum=s}}\n\n====Derived terms====\n{{col-auto|pl|title=adjective|montrealski}}\n\n===Further reading===\n* {{R:pl:PWN}}\n\n==Portuguese==\n\n===Alternative forms===\n* {{alt|pt|Montréal}}\n\n===Etymology===\nBorrowed from {{bor|pt|fr|Montréal}}.\n\n===Pronunciation===\n{{pt-IPA}}\n\n===Proper noun===\n{{pt-proper noun|f}}\n\n# {{place|pt|largest city|p/Quebec|c/Canada|t1=Montreal}}\n\n==Slovak==\n\n===Pronunciation===\n* {{sk-IPA}}\n\n===Proper noun===\n{{sk-proper noun|m-in|gen=Montrealu|decl=dub}}\n\n# {{place|sk|largest city|p/Quebec|c/Canada|t=Montreal}}\n\n===References===\n* {{R:sk:SDK}}\n\n==Spanish==\n\n===Alternative forms===\n* {{alt|es|Montréal}}\n\n===Pronunciation===\n{{es-pr}}\n\n===Proper noun===\n{{es-proper noun|f}}\n\n# {{place|es|largest city|p/Quebec|c/Canada|t1=Montreal}}\n\n==Turkish==\n\n===Etymology===\nFrom {{der|tr|fr|Montréal}}\n\n===Pronunciation===\n* {{IPA|tr|/mont.ɾeˈɑl/}}\n\n===Noun===\n{{tr-noun|i|ler}}\n\n# [[Montreal]]\n\n====Declension====\n{{tr-infl-noun-c|e|poss=1}}"
        },
        "license": [
            {
                "name": "Creative Commons Attribution Share Alike 3.0 Unported",
                "identifier": "CC-BY-SA-3.0",
                "url": "https://creativecommons.org/licenses/by-sa/3.0/"
            }
        ],
        "event": {
            "identifier": "20656eda-86bf-4d74-9817-25f52d1edb69",
            "type": "update",
            "date_created": "2024-02-29T04:54:55.140642Z",
            "date_published": "2024-02-29T04:54:59.151524Z"
        }
    },
    {
        "name": "Montreal",
        "identifier": 7954681,
        "abstract": "Montreal is the second most populous city in Canada, the tenth most populous city in North America, and the most populous city in the province of Quebec. Founded in 1642 as Ville-Marie, or \"City of Mary\", it is named after Mount Royal, the triple-peaked hill around which the early city of Ville-Marie was built. The city is centred on the Island of Montreal, which obtained its name from the same origin as the city, and a few much smaller peripheral islands, the largest of which is Île Bizard. The city is 196 km (122 mi) east of the national capital, Ottawa, and 258 km (160 mi) southwest of the provincial capital, Quebec City. As of 2021, the city has a population of 1,762,949, and a metropolitan population of 4,291,732, making it the second-largest city, and second-largest metropolitan area in Canada. French is the city's official language. In 2021, 85.7% of the population of the city of Montreal considered themselves fluent in French while 90.2% could speak it in the metropolitan area. Montreal is one of the most bilingual cities in Quebec and Canada, with 58.5% of the population able to speak both English and French. Historically the commercial capital of Canada, Montreal was surpassed in population and economic strength by Toronto in the 1970s. Montreal remains an important centre of art, culture, literature, film and television, music, commerce, aerospace, transport, finance, pharmaceuticals, technology, design, education, tourism, food, fashion, video game development, and world affairs. Montreal is the location of the headquarters of the International Civil Aviation Organization, and was named a UNESCO City of Design in 2006. In 2017, Montreal was ranked the 12th-most liveable city in the world by the Economist Intelligence Unit in its annual Global Liveability Ranking, although it slipped to rank 40 in the 2021 index, primarily due to stress on the healthcare system from the COVID-19 pandemic. It is regularly ranked as a top ten city in the world to be a university student in the QS World University Rankings. Montreal has hosted multiple international conferences and events, including the 1967 International and Universal Exposition and the 1976 Summer Olympics. It is the only Canadian city to have held the Summer Olympics. In 2018, Montreal was ranked as a global city. The city hosts the Canadian Grand Prix of Formula One; the Montreal International Jazz Festival, the largest jazz festival in the world; the Just for Laughs festival, the largest comedy festival in the world; and Les Francos de Montréal, the largest French-language music festival in the world. In sports, it is home to the Montreal Canadiens of the National Hockey League, who have won the Stanley Cup more times than any other team.",
        "version": {
            "identifier": 1211753567,
            "tags": [
                "wikieditor"
            ],
            "scores": {
                "revertrisk": {
                    "prediction": true,
                    "probability": {
                        "false": 0.015385448932647705,
                        "true": 0.9846145510673523
                    }
                }
            },
            "editor": {
                "name": "70.55.131.207",
                "groups": [
                    "*"
                ],
                "is_anonymous": true,
                "date_started": "0001-01-01T00:00:00Z"
            },
            "number_of_characters": 212346,
            "size": {
                "value": 213010,
                "unit_text": "B"
            },
            "maintenance_tags": {
                "citation_needed_count": 15,
                "clarification_needed_count": 1,
                "update_count": 1
            }
        },
        "url": "https://en.wikipedia.org/wiki/Montreal",
        "date_created": "2001-10-13T18:26:50Z",
        "date_modified": "2024-03-04T07:14:40Z",
        "main_entity": {
            "identifier": "Q340",
            "url": "https://www.wikidata.org/entity/Q340"
        },
        "is_part_of": {
            "identifier": "enwiki",
            "url": "https://en.wikipedia.org"
        },
        "additional_entities": [
            {
                "identifier": "Q150",
                "url": "https://www.wikidata.org/entity/Q150",
                "aspects": [
                    "O"
                ]
            },
            {
                "identifier": "Q1860",
                "url": "https://www.wikidata.org/entity/Q1860",
                "aspects": [
                    "O"
                ]
            },
            {
                "identifier": "Q340",
                "url": "https://www.wikidata.org/entity/Q340",
                "aspects": [
                    "C",
                    "D.en",
                    "O",
                    "S",
                    "T"
                ]
            }
        ],
        "in_language": {
            "identifier": "en"
        },
        "image": {
            "content_url": "https://upload.wikimedia.org/wikipedia/commons/d/dc/Flag_of_Montreal.svg",
            "width": 435,
            "height": 217
        },
        "description": "Largest city in Quebec, Canada",
        "infobox": [
            {
                "name": "Infobox settlement\n",
                "type": "infobox",
                "has_parts": [
                    {
                        "name": "Montreal Montréal (French)",
                        "type": "section"
                    },
                    {
                        "name": "City",
                        "type": "section",
                        "has_parts": [
                            {
                                "type": "field",
                                "value": "Ville de Montréal"
                            },
                            {
                                "type": "field",
                                "value": "From top, left to right: Downtown Montreal skyline, Old Montreal, Notre-Dame Basilica, Old Port of Montreal, Saint Joseph's Oratory, Olympic Stadium",
                                "images": [
                                    {
                                        "content_url": "https://upload.wikimedia.org/wikipedia/commons/thumb/9/9f/Montreal_Montage_2020.jpg/258px-Montreal_Montage_2020.jpg",
                                        "height": 258,
                                        "width": 258
                                    }
                                ],
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Downtown_Montreal",
                                        "text": "Downtown Montreal"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Old_Montreal",
                                        "text": "Old Montreal"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Notre-Dame_Basilica_(Montreal)",
                                        "text": "Notre-Dame Basilica"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Old_Port_of_Montreal",
                                        "text": "Old Port of Montreal"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Saint_Joseph's_Oratory",
                                        "text": "Saint Joseph's Oratory"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Olympic_Stadium_(Montreal)",
                                        "text": "Olympic Stadium"
                                    }
                                ]
                            },
                            {
                                "type": "field",
                                "value": "Flag Coat of arms Logo",
                                "images": [
                                    {
                                        "content_url": "https://upload.wikimedia.org/wikipedia/commons/thumb/d/dc/Flag_of_Montreal.svg/125px-Flag_of_Montreal.svg.png",
                                        "alternative_text": "Flag of Montreal",
                                        "height": 125,
                                        "width": 125
                                    },
                                    {
                                        "content_url": "https://upload.wikimedia.org/wikipedia/commons/thumb/4/47/Coat_of_arms_of_Montreal.svg/95px-Coat_of_arms_of_Montreal.svg.png",
                                        "alternative_text": "Official seal of Montreal",
                                        "height": 95,
                                        "width": 95
                                    },
                                    {
                                        "content_url": "https://upload.wikimedia.org/wikipedia/en/thumb/7/7e/City_of_Montr%C3%A9al_logo.svg/160px-City_of_Montr%C3%A9al_logo.svg.png",
                                        "alternative_text": "Official logo of Montreal",
                                        "height": 160,
                                        "width": 160
                                    }
                                ],
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Flag_of_Montreal",
                                        "text": "Flag"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Coat_of_arms_of_Montreal",
                                        "text": "Coat of arms"
                                    }
                                ]
                            },
                            {
                                "type": "field",
                                "value": "Nickname(s): \"MTL\", \"The 514\", \"The City of Festivals\", \"The City of Saints\", \"The City of a Hundred Steeples\", \"Sin City\", \"La Métropole\"",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Name_of_Montreal#Nicknames",
                                        "text": "\"MTL\", \"The 514\", \"The City of Festivals\", \"The City of Saints\", \"The City of a Hundred Steeples\", \"Sin City\", \"La Métropole\""
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Montreal#cite_note-1"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Montreal#cite_note-2"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Montreal#cite_note-3"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Montreal#cite_note-4"
                                    }
                                ]
                            },
                            {
                                "type": "field",
                                "value": "Motto: Concordia Salus (\"well-being through harmony\")",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Concordia_Salus",
                                        "text": "Concordia Salus"
                                    }
                                ]
                            },
                            {
                                "type": "field",
                                "value": "Interactive map of Montreal",
                                "images": [
                                    {
                                        "content_url": "https:https://maps.wikimedia.org/img/osm-intl,4,a,a,300x200.png?lang=en&domain=en.wikipedia.org&title=Montreal&revid=1209145403&groups=_859f4c24fa1975fa70eb5d146af4bfd627adbc0d",
                                        "height": 300,
                                        "width": 300
                                    }
                                ],
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/",
                                        "images": [
                                            {
                                                "content_url": "https:https://maps.wikimedia.org/img/osm-intl,4,a,a,300x200.png?lang=en&domain=en.wikipedia.org&title=Montreal&revid=1209145403&groups=_859f4c24fa1975fa70eb5d146af4bfd627adbc0d",
                                                "height": 300,
                                                "width": 300
                                            }
                                        ]
                                    }
                                ]
                            },
                            {
                                "type": "field",
                                "value": "Coordinates: 45°30′32″N 73°33′15″W / 45.50889°N 73.55417°W",
                                "links": [
                                    {
                                        "url": "https://geohack.toolforge.org/geohack.php?pagename=Montreal&params=45_30_32_N_73_33_15_W_region:CA-QC_type:city(1,800,000)",
                                        "text": "45°30′32″N 73°33′15″W / 45.50889°N 73.55417°W"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Montreal#cite_note-5"
                                    }
                                ]
                            },
                            {
                                "name": "Country",
                                "type": "field",
                                "value": "Canada"
                            },
                            {
                                "name": "Province",
                                "type": "field",
                                "value": "Quebec",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Provinces_and_Territories_of_Canada",
                                        "text": "Province"
                                    }
                                ]
                            },
                            {
                                "name": "Region",
                                "type": "field",
                                "value": "Montreal",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Region_(Quebec)",
                                        "text": "Region"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Montreal_(region)",
                                        "text": "Montreal"
                                    }
                                ]
                            },
                            {
                                "name": "Urban agglomeration",
                                "type": "field",
                                "value": "Montreal",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Urban_agglomerations_in_Quebec",
                                        "text": "Urban agglomeration"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Urban_agglomeration_of_Montreal",
                                        "text": "Montreal"
                                    }
                                ]
                            },
                            {
                                "name": "Founded",
                                "type": "field",
                                "value": "May 17, 1642"
                            },
                            {
                                "name": "Incorporated",
                                "type": "field",
                                "value": "1832"
                            },
                            {
                                "name": "Constituted",
                                "type": "field",
                                "value": "January 1, 2002"
                            },
                            {
                                "name": "Named for",
                                "type": "field",
                                "value": "Mount Royal",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Namesake",
                                        "text": "Named for"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Mount_Royal",
                                        "text": "Mount Royal"
                                    }
                                ]
                            },
                            {
                                "name": "Boroughs",
                                "type": "field",
                                "value": "List Ahuntsic-Cartierville Anjou Côte-des-Neiges–Notre-Dame-de-Grâce L'Île-Bizard–Sainte-Geneviève LaSalle Lachine Le Plateau-Mont-Royal Le Sud-Ouest Mercier–Hochelaga-Maisonneuve Montréal-Nord Outremont Pierrefonds-Roxboro Rivière-des-Prairies–Pointe-aux-Trembles Rosemont–La Petite-Patrie Saint-Laurent Saint-Léonard Verdun Ville-Marie Villeray–Saint-Michel–Parc-Extension",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Ahuntsic-Cartierville",
                                        "text": "Ahuntsic-Cartierville"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Anjou,_Quebec",
                                        "text": "Anjou"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Côte-des-Neiges–Notre-Dame-de-Grâce",
                                        "text": "Côte-des-Neiges–Notre-Dame-de-Grâce"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/L'Île-Bizard–Sainte-Geneviève",
                                        "text": "L'Île-Bizard–Sainte-Geneviève"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/LaSalle,_Quebec",
                                        "text": "LaSalle"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Lachine,_Quebec",
                                        "text": "Lachine"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Le_Plateau-Mont-Royal",
                                        "text": "Le Plateau-Mont-Royal"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Le_Sud-Ouest",
                                        "text": "Le Sud-Ouest"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Mercier–Hochelaga-Maisonneuve",
                                        "text": "Mercier–Hochelaga-Maisonneuve"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Montréal-Nord",
                                        "text": "Montréal-Nord"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Outremont,_Quebec",
                                        "text": "Outremont"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Pierrefonds-Roxboro",
                                        "text": "Pierrefonds-Roxboro"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Rivière-des-Prairies–Pointe-aux-Trembles",
                                        "text": "Rivière-des-Prairies–Pointe-aux-Trembles"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Rosemont–La_Petite-Patrie",
                                        "text": "Rosemont–La Petite-Patrie"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Saint-Laurent,_Quebec",
                                        "text": "Saint-Laurent"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/St._Leonard,_Quebec",
                                        "text": "Saint-Léonard"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Verdun,_Quebec",
                                        "text": "Verdun"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Ville-Marie,_Montreal",
                                        "text": "Ville-Marie"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Villeray–Saint-Michel–Parc-Extension",
                                        "text": "Villeray–Saint-Michel–Parc-Extension"
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "name": "Government",
                        "type": "section",
                        "has_parts": [
                            {
                                "name": "Type",
                                "type": "field",
                                "value": "Montreal City Council",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Montreal_City_Council",
                                        "text": "Montreal City Council"
                                    }
                                ]
                            },
                            {
                                "name": "Mayor",
                                "type": "field",
                                "value": "Valérie Plante",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Valérie_Plante",
                                        "text": "Valérie Plante"
                                    }
                                ]
                            },
                            {
                                "name": "Federal riding",
                                "type": "field",
                                "value": "List",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/List_of_Canadian_federal_electoral_districts",
                                        "text": "Federal riding"
                                    }
                                ]
                            },
                            {
                                "name": "Provincial riding",
                                "type": "field",
                                "value": "List",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/List_of_Quebec_provincial_electoral_districts",
                                        "text": "Provincial riding"
                                    }
                                ]
                            },
                            {
                                "name": "MPs",
                                "type": "field",
                                "value": "List of MPs",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Member_of_Parliament_(Canada)",
                                        "text": "MPs"
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "name": "Area",
                        "type": "section",
                        "has_parts": [
                            {
                                "name": "City",
                                "type": "field",
                                "value": "431.50 km2 (166.60 sq mi)",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/List_of_cities_in_Quebec",
                                        "text": "City"
                                    }
                                ]
                            },
                            {
                                "name": "Land",
                                "type": "field",
                                "value": "365.13 km2 (140.98 sq mi)"
                            },
                            {
                                "name": "Urban",
                                "type": "field",
                                "value": "1,293.99 km2 (499.61 sq mi)",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Montreal#cite_note-cp2011-PC-10"
                                    }
                                ]
                            },
                            {
                                "name": "Metro",
                                "type": "field",
                                "value": "4,604.26 km2 (1,777.71 sq mi)",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Montreal#cite_note-cp2011-CA-11"
                                    }
                                ]
                            },
                            {
                                "name": "Highest elevation",
                                "type": "field",
                                "value": "233 m (764 ft)"
                            },
                            {
                                "name": "Lowest elevation",
                                "type": "field",
                                "value": "6 m (20 ft)"
                            }
                        ]
                    },
                    {
                        "name": "Population (2021)",
                        "type": "section",
                        "has_parts": [
                            {
                                "name": "City",
                                "type": "field",
                                "value": "1,762,949 (2nd)",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/List_of_cities_in_Quebec",
                                        "text": "City"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/List_of_the_largest_population_centres_in_Canada",
                                        "text": "2nd"
                                    }
                                ]
                            },
                            {
                                "name": "Density",
                                "type": "field",
                                "value": "4,828.3/km2 (12,505/sq mi)"
                            },
                            {
                                "name": "Metro",
                                "type": "field",
                                "value": "4,291,732 (2nd)",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Metropolitan_area",
                                        "text": "Metro"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Montreal#cite_note-cp2016-CA-12"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/List_of_census_metropolitan_areas_and_agglomerations_in_Canada",
                                        "text": "2nd"
                                    }
                                ]
                            },
                            {
                                "name": "Metro density",
                                "type": "field",
                                "value": "919/km2 (2,380/sq mi)"
                            },
                            {
                                "name": "Demonym(s)",
                                "type": "field",
                                "value": "Montrealer Montréalais(e)",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Demonym",
                                        "text": "Demonym(s)"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Montreal#cite_note-13"
                                    }
                                ]
                            },
                            {
                                "name": "Time zone",
                                "type": "field",
                                "value": "UTC−05:00 (EST)",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Time_zone",
                                        "text": "Time zone"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/UTC−05:00",
                                        "text": "UTC−05:00"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Eastern_Time_Zone",
                                        "text": "EST"
                                    }
                                ]
                            },
                            {
                                "name": "Summer (DST)",
                                "type": "field",
                                "value": "UTC−04:00 (EDT)",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Daylight_saving_time",
                                        "text": "DST"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/UTC−04:00",
                                        "text": "UTC−04:00"
                                    }
                                ]
                            },
                            {
                                "name": "Postal codes",
                                "type": "field",
                                "value": "H H1A, H1C-H3N, H3S-H3W, H4A-H4T, H4Z-H5B, H8R-H8Z, H9C-H9E, H9H, H9K",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Canadian_postal_code",
                                        "text": "Postal codes"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/List_of_H_postal_codes_of_Canada",
                                        "text": "H"
                                    }
                                ]
                            },
                            {
                                "name": "Area code(s)",
                                "type": "field",
                                "value": "514 and 438 and 263",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Telephone_numbering_plan",
                                        "text": "Area code(s)"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Area_codes_514_and_438",
                                        "text": "514 and 438"
                                    }
                                ]
                            },
                            {
                                "name": "Police",
                                "type": "field",
                                "value": "SPVM",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Service_de_police_de_la_Ville_de_Montréal",
                                        "text": "SPVM"
                                    }
                                ]
                            },
                            {
                                "name": "GDP (Montreal CMA)",
                                "type": "field",
                                "value": "C$221.9 billion (2018)",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/GDP",
                                        "text": "GDP"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Montreal#cite_note-14"
                                    }
                                ]
                            },
                            {
                                "name": "GDP per capita (Montreal CMA)",
                                "type": "field",
                                "value": "C$48,289 (2022)",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Montreal#cite_note-15"
                                    }
                                ]
                            },
                            {
                                "name": "Website",
                                "type": "field",
                                "value": "montreal .ca /en /",
                                "links": [
                                    {
                                        "url": "https://montreal.ca/en/",
                                        "text": "montreal .ca /en /"
                                    }
                                ]
                            }
                        ]
                    }
                ]
            }
        ],
        "article_sections": [
            {
                "name": "Abstract",
                "type": "section",
                "has_parts": [
                    {
                        "type": "paragraph",
                        "value": "Montreal (CA: /ˌ m ʌ n t r i ˈ ɔː l / MUN -tree-AWL; French: Montréal) is the second most populous city in Canada, the tenth most populous city in North America, and the most populous city in the province of Quebec . Founded in 1642 as Ville-Marie, or \"City of Mary\", it is named after Mount Royal, the triple-peaked hill around which the early city of Ville-Marie was built. The city is centred on the Island of Montreal, which obtained its name from the same origin as the city, and a few much smaller peripheral islands, the largest of which is Île Bizard . The city is 196 km (122 mi) east of the national capital, Ottawa, and 258 km (160 mi) southwest of the provincial capital, Quebec City .",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Canadian_English",
                                "text": "CA"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Help:IPA/English",
                                "text": "/ˌ m ʌ n t r i ˈ ɔː l /"
                            },
                            {
                                "url": "https://upload.wikimedia.org/wikipedia/commons/transcoded/5/57/Montreal-English-pronunciation.oga/Montreal-English-pronunciation.oga.mp3",
                                "text": "Play audio"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Help:Pronunciation_respelling_key",
                                "text": "MUN -tree-AWL"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/French_language",
                                "text": "French"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Help:IPA/French",
                                "text": "Help:IPA/French"
                            },
                            {
                                "url": "https://upload.wikimedia.org/wikipedia/commons/transcoded/b/b6/Qc-Montr%C3%A9al.ogg/Qc-Montr%C3%A9al.ogg.mp3",
                                "text": "Play audio"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/List_of_the_largest_municipalities_in_Canada_by_population",
                                "text": "second most populous city"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Canada",
                                "text": "Canada"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/List_of_North_American_cities_by_population",
                                "text": "tenth most populous city"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/North_America",
                                "text": "North America"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/List_of_towns_in_Quebec",
                                "text": "most populous city"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Provinces_and_territories_of_Canada",
                                "text": "province"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Quebec",
                                "text": "Quebec"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Fort_Ville-Marie",
                                "text": "Ville-Marie"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Mount_Royal",
                                "text": "Mount Royal"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Island_of_Montreal",
                                "text": "Island of Montreal"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Île_Bizard",
                                "text": "Île Bizard"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Ottawa",
                                "text": "Ottawa"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Quebec_City",
                                "text": "Quebec City"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "As of 2021, the city has a population of 1,762,949, and a metropolitan population of 4,291,732, making it the second-largest city, and second-largest metropolitan area in Canada. French is the city's official language. In 2021, 85.7% of the population of the city of Montreal considered themselves fluent in French while 90.2% could speak it in the metropolitan area. Montreal is one of the most bilingual cities in Quebec and Canada, with 58.5% of the population able to speak both English and French.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Census_Metropolitan_Area#Census_metropolitan_areas",
                                "text": "metropolitan"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/List_of_the_largest_municipalities_in_Canada_by_population",
                                "text": "second-largest"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/List_of_census_metropolitan_areas_and_agglomerations_in_Canada",
                                "text": "second-largest"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/French_language",
                                "text": "French"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Bilingualism",
                                "text": "bilingual"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "Historically the commercial capital of Canada, Montreal was surpassed in population and economic strength by Toronto in the 1970s. Montreal remains an important centre of art, culture, literature, film and television, music, commerce, aerospace, transport, finance, pharmaceuticals, technology, design, education, tourism, food, fashion, video game development, and world affairs. Montreal is the location of the headquarters of the International Civil Aviation Organization, and was named a UNESCO City of Design in 2006. In 2017, Montreal was ranked the 12th-most liveable city in the world by the Economist Intelligence Unit in its annual Global Liveability Ranking, although it slipped to rank 40 in the 2021 index, primarily due to stress on the healthcare system from the COVID-19 pandemic . It is regularly ranked as a top ten city in the world to be a university student in the QS World University Rankings .",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Toronto",
                                "text": "Toronto"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Culture_of_Montreal",
                                "text": "culture"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/List_of_films_shot_in_Montreal",
                                "text": "film"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Transportation_in_Montreal",
                                "text": "transport"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Education_in_Montreal",
                                "text": "education"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Tourism_in_Montreal",
                                "text": "tourism"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/International_Civil_Aviation_Organization",
                                "text": "International Civil Aviation Organization"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/UNESCO",
                                "text": "UNESCO"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Design_Cities_(UNESCO)",
                                "text": "City of Design"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Economist_Intelligence_Unit",
                                "text": "Economist Intelligence Unit"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Global_Liveability_Ranking",
                                "text": "Global Liveability Ranking"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/COVID-19_pandemic",
                                "text": "COVID-19 pandemic"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/QS_World_University_Rankings",
                                "text": "QS World University Rankings"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "Montreal has hosted multiple international conferences and events, including the 1967 International and Universal Exposition and the 1976 Summer Olympics . It is the only Canadian city to have held the Summer Olympics. In 2018, Montreal was ranked as a global city . The city hosts the Canadian Grand Prix of Formula One; the Montreal International Jazz Festival, the largest jazz festival in the world; the Just for Laughs festival, the largest comedy festival in the world; and Les Francos de Montréal, the largest French-language music festival in the world. In sports, it is home to the Montreal Canadiens of the National Hockey League, who have won the Stanley Cup more times than any other team.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Expo_67",
                                "text": "1967 International and Universal Exposition"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/1976_Summer_Olympics",
                                "text": "1976 Summer Olympics"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Globalization_and_World_Cities_Research_Network",
                                "text": "global city"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Canadian_Grand_Prix",
                                "text": "Canadian Grand Prix"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Formula_One",
                                "text": "Formula One"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Montreal_International_Jazz_Festival",
                                "text": "Montreal International Jazz Festival"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Just_for_Laughs",
                                "text": "Just for Laughs"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Les_Francos_de_Montréal",
                                "text": "Les Francos de Montréal"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Sports_in_Montreal",
                                "text": "sports"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Montreal_Canadiens",
                                "text": "Montreal Canadiens"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/National_Hockey_League",
                                "text": "National Hockey League"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Stanley_Cup",
                                "text": "Stanley Cup"
                            }
                        ]
                    }
                ]
            },
            {
                "name": "Etymology and original names",
                "type": "section",
                "has_parts": [
                    {
                        "type": "paragraph",
                        "value": "In the Ojibwe language, the land is called Mooniyaang which was \"the first stopping place\" in the Ojibwe migration story as related in the seven fires prophecy .",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Ojibwe_language",
                                "text": "Ojibwe language"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Seven_fires_prophecy",
                                "text": "seven fires prophecy"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "In the Mohawk language, the land is called Tiohtià:ke . This is an abbreviation of Teionihtiohtiá:kon, which loosely translates as \"where the group divided/parted ways.\"",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Mohawk_language",
                                "text": "Mohawk language"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "French settlers from La Flèche in the Loire valley first named their new town, founded in 1642, Ville Marie (\"City of Mary\"), named for the Virgin Mary .",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/La_Flèche",
                                "text": "La Flèche"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Fort_Ville-Marie",
                                "text": "Ville Marie"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Virgin_Mary",
                                "text": "Virgin Mary"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "The current form of the name, Montréal, is generally thought to be derived from Mount Royal (Mont Royal in French), the triple-peaked hill in the heart of the city. There are multiple explanations for how Mont Royal became Montréal . In 16th century French, the forms réal and royal were used interchangeably, so Montréal could simply be a variant of Mont Royal . In the second explanation, the name came from an Italian translation. Venetian geographer Giovanni Battista Ramusio used the name Monte Real to designate Mount Royal in his 1556 map of the region. However, the Commission de toponymie du Québec disputes this explanation.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Giovanni_Battista_Ramusio",
                                "text": "Giovanni Battista Ramusio"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Commission_de_toponymie_du_Québec",
                                "text": "Commission de toponymie du Québec"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "Historiographer François de Belleforest was the first to use the form Montréal with reference to the entire region in 1575."
                    }
                ]
            },
            {
                "name": "History",
                "type": "section",
                "has_parts": [
                    {
                        "name": "Pre-European contact",
                        "type": "section",
                        "has_parts": [
                            {
                                "type": "paragraph",
                                "value": "Archaeological evidence in the region indicate that First Nations native people occupied the island of Montreal as early as 4,000 years ago. By the year AD 1000, they had started to cultivate maize . Within a few hundred years, they had built fortified villages. The Saint Lawrence Iroquoians, an ethnically and culturally distinct group from the Iroquois nations of the Haudenosaunee (then based in present-day New York), established the village of Hochelaga at the foot of Mount Royal two centuries before the French arrived. Archeologists have found evidence of their habitation there and at other locations in the valley since at least the 14th century.The French explorer Jacques Cartier visited Hochelaga on October 2, 1535, and estimated the population of the native people at Hochelaga to be \"over a thousand people\". Evidence of earlier occupation of the island, such as those uncovered in 1642 during the construction of Fort Ville-Marie, have effectively been removed.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Archaeological",
                                        "text": "Archaeological"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/First_Nations_in_Canada",
                                        "text": "First Nations"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Maize",
                                        "text": "maize"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Fortified",
                                        "text": "fortified"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Saint_Lawrence_Iroquoians",
                                        "text": "Saint Lawrence Iroquoians"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Iroquois",
                                        "text": "Iroquois"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Haudenosaunee",
                                        "text": "Haudenosaunee"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Hochelaga_(village)",
                                        "text": "village of Hochelaga"
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "name": "Early European settlement (1600–1760)",
                        "type": "section",
                        "has_parts": [
                            {
                                "type": "paragraph",
                                "value": "In 1603, French explorer Samuel de Champlain reported that the St Lawrence Iroquoians and their settlements had disappeared altogether from the St Lawrence valley. This is believed to be due to outmigration, epidemics of European diseases, or intertribal wars. In 1611, Champlain established a fur trading post on the Island of Montreal on a site initially named La Place Royale . At the confluence of Petite Riviere and St. Lawrence River, it is where present-day Pointe-à-Callière stands. On his 1616 map, Champlain named the island Lille de Villemenon in honour of the sieur de Villemenon, a French dignitary who was seeking the viceroyship of New France. In 1639, Jérôme Le Royer de La Dauversière obtained the Seigneurial title to the Island of Montreal in the name of the Notre Dame Society of Montreal to establish a Roman Catholic mission to evangelize natives.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Samuel_de_Champlain",
                                        "text": "Samuel de Champlain"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Fur",
                                        "text": "fur"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Trading_post",
                                        "text": "trading post"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Pointe-à-Callière",
                                        "text": "Pointe-à-Callière"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Jérôme_Le_Royer_de_La_Dauversière",
                                        "text": "Jérôme Le Royer de La Dauversière"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Seigneurial_system_of_New_France",
                                        "text": "Seigneurial title"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Société_de_Notre-Dame_de_Montréal",
                                        "text": "Notre Dame Society of Montreal"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Mission_(Christian)",
                                        "text": "mission"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Evangelize",
                                        "text": "evangelize"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "Dauversiere hired Paul Chomedey de Maisonneuve, then age 30, to lead a group of colonists to build a mission on his new seigneury. The colonists left France in 1641 for Quebec and arrived on the island the following year. On May 17, 1642, Ville-Marie was founded on the southern shore of Montreal island, with Maisonneuve as its first governor. The settlement included a chapel and a hospital, under the command of Jeanne Mance . By 1643, Ville-Marie had come under Iroquois raids. In 1652, Maisonneuve returned to France to raise 100 volunteers to bolster the colonial population. If the effort had failed, Montreal was to be abandoned and the survivors re-located downriver to Quebec City . Before these 100 arrived in the fall of 1653, the population of Montreal was barely 50 people.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Jérome_le_Royer_de_la_Dauversiere",
                                        "text": "Dauversiere"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Paul_Chomedey_de_Maisonneuve",
                                        "text": "Paul Chomedey de Maisonneuve"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Jeanne_Mance",
                                        "text": "Jeanne Mance"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Quebec_City",
                                        "text": "Quebec City"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "By 1685, Ville-Marie was home to some 600 colonists, most of them living in modest wooden houses. Ville-Marie became a centre for the fur trade and a base for further exploration . In 1689, the English-allied Iroquois attacked Lachine on the Island of Montreal, committing the worst massacre in the history of New France. By the early 18th century, the Sulpician Order was established there. To encourage French settlement, it wanted the Mohawk to move away from the fur trading post at Ville-Marie. It had a mission village, known as Kahnewake, south of the St Lawrence River. The fathers persuaded some Mohawk to make a new settlement at their former hunting grounds north of the Ottawa River. This became Kanesatake . In 1745, several Mohawk families moved upriver to create another settlement, known as Akwesasne . All three are now Mohawk reserves in Canada. The Canadian territory was ruled as a French colony until 1760, when Montreal fell to a British offensive during the Seven Years' War . The colony then surrendered to Great Britain.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Fur_trade",
                                        "text": "fur trade"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/French_colonization_of_the_Americas",
                                        "text": "exploration"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Lachine_massacre",
                                        "text": "Lachine"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Sulpician_Order",
                                        "text": "Sulpician Order"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Kahnewake",
                                        "text": "Kahnewake"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Kanesatake",
                                        "text": "Kanesatake"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Akwesasne",
                                        "text": "Akwesasne"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Montreal_Campaign",
                                        "text": "Montreal fell to a British offensive"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Seven_Years'_War",
                                        "text": "Seven Years' War"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "Ville-Marie was the name for the settlement that appeared in all official documents until 1705, when Montreal appeared for the first time, although people referred to the \"Island of Montreal\" long before then."
                            }
                        ]
                    },
                    {
                        "name": "American occupation (1775–1776)",
                        "type": "section",
                        "has_parts": [
                            {
                                "type": "paragraph",
                                "value": "As part of the American Revolution, the invasion of Quebec resulted after Benedict Arnold captured Fort Ticonderoga in present-day upstate New York in May 1775 as a launching point to Arnold's invasion of Quebec in September . While Arnold approached the Plains of Abraham, Montreal fell to American forces led by Richard Montgomery on November 13, 1775, after it was abandoned by Guy Carleton . After Arnold withdrew from Quebec City to Pointe-aux-Trembles on November 19, Montgomery's forces left Montreal on December 1 and arrived there on December 3 to plot to attack Quebec City, with Montgomery leaving David Wooster in charge of the city. Montgomery was killed in the failed attack and Arnold, who had taken command, sent Brigadier General Moses Hazen to inform Wooster of the defeat.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/American_Revolution",
                                        "text": "American Revolution"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Invasion_of_Quebec_(1775)",
                                        "text": "invasion of Quebec"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Benedict_Arnold",
                                        "text": "Benedict Arnold"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Fort_Ticonderoga",
                                        "text": "Fort Ticonderoga"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Benedict_Arnold's_expedition_to_Quebec",
                                        "text": "Arnold's invasion of Quebec in September"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Plains_of_Abraham",
                                        "text": "Plains of Abraham"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Richard_Montgomery",
                                        "text": "Richard Montgomery"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Guy_Carleton,_1st_Baron_Dorchester",
                                        "text": "Guy Carleton"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Pointe-aux-Trembles",
                                        "text": "Pointe-aux-Trembles"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Battle_of_Quebec_(1775)",
                                        "text": "attack Quebec City"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/David_Wooster",
                                        "text": "David Wooster"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Moses_Hazen",
                                        "text": "Moses Hazen"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "Wooster left Hazen in command on March 20, 1776, as he left to replace Arnold in leading further attacks on Quebec City. On April 19, Arnold arrived in Montreal to take over command from Hazen, who remained as his second-in-command. Hazen sent Colonel Timothy Bedel to form a garrison of 390 men 40 miles upriver in a garrison at Les Cèdres, Quebec, to defend Montreal against the British army. In the Battle of the Cedars, Bedel's lieutenant Isaac Butterfield surrendered to George Forster.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Timothy_Bedel",
                                        "text": "Timothy Bedel"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Les_Cèdres,_Quebec",
                                        "text": "Les Cèdres, Quebec"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Battle_of_the_Cedars",
                                        "text": "Battle of the Cedars"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Isaac_Butterfield",
                                        "text": "Isaac Butterfield"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "Forster advanced to Fort Senneville on May 23. By May 24, Arnold was entrenched in Montreal's borough of Lachine . Forster initially approached Lachine, then withdrew to Quinze-Chênes . Arnold's forces then abandoned Lachine to chase Forster. The Americans burned Senneville on May 26. After Arnold crossed the Ottawa River in pursuit of Forster, Forster's cannons repelled Arnold's forces. Forster negotiated a prisoner exchange with Henry Sherburne and Isaac Butterfield, resulting in a May 27 boating of their deputy Lieutenant Park being returned to the Americans. Arnold and Forster negotiated further and more American prisoners were returned to Arnold at Sainte-Anne-de-Bellevue, Quebec, (\"Fort Anne\") on May 30 (delayed two days by wind).",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Fort_Senneville",
                                        "text": "Fort Senneville"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Lachine,_Quebec",
                                        "text": "Montreal's borough of Lachine"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Quinze-Chênes?action=edit&redlink=1",
                                        "text": "Quinze-Chênes"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Ottawa_River",
                                        "text": "Ottawa River"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Henry_Sherburne",
                                        "text": "Henry Sherburne"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Sainte-Anne-de-Bellevue,_Quebec",
                                        "text": "Sainte-Anne-de-Bellevue, Quebec"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "Arnold eventually withdrew his forces back to the New York fort of Ticonderoga by the summer. On June 15, Arnold's messenger approaching Sorel spotted Carleton returning with a fleet of ships and notified him. Arnold's forces abandoned Montreal (attempting to burn it down in the process) prior to the June 17 arrival of Carleton's fleet.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Sorel-Tracy",
                                        "text": "Sorel"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "The Americans did not return British prisoners in exchange, as previously agreed, due to accusations of abuse, with Congress repudiating the agreement at the protest of George Washington. Arnold blamed Colonel Timothy Bedel for the defeat, removing him and Lieutenant Butterfield from command and sending them to Sorel for court-martial. The retreat of the American army delayed their court martial until August 1, 1776, when they were convicted and cashiered at Ticonderoga. Bedel was given a new commission by Congress in October 1777 after Arnold was assigned to defend Rhode Island in July 1777 .",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Cashiered",
                                        "text": "cashiered"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Military_career_of_Benedict_Arnold,_1777–1779",
                                        "text": "July 1777"
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "name": "Modern history as city (1832–present)",
                        "type": "section",
                        "has_parts": [
                            {
                                "type": "paragraph",
                                "value": "Montreal was incorporated as a city in 1832. The opening of the Lachine Canal permitted ships to bypass the unnavigable Lachine Rapids, while the construction of the Victoria Bridge established Montreal as a major railway hub. The leaders of Montreal's business community had started to build their homes in the Golden Square Mile from about 1850. By 1860, it was the largest municipality in British North America and the undisputed economic and cultural centre of Canada.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Lachine_Rapids",
                                        "text": "Lachine Rapids"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Victoria_Bridge_(Montreal)",
                                        "text": "Victoria Bridge"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Golden_Square_Mile",
                                        "text": "Golden Square Mile"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/British_North_America",
                                        "text": "British North America"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "In the 19th century, maintaining Montreal's drinking water became increasingly difficult with the rapid increase in population. A majority of the drinking water was still coming from the city's harbour, which was busy and heavily trafficked, leading to the deterioration of the water within. In the mid-1840s, the City of Montreal installed a water system that would pump water from the St. Lawrence and into cisterns . The cisterns would then be transported to the desired location. This was not the first water system of its type in Montreal, as there had been one in private ownership since 1801. In the middle of the 19th century, water distribution was carried out by \"fontainiers\". The fountainiers would open and close water valves outside of buildings, as directed, all over the city. As they lacked modern plumbing systems it was impossible to connect all buildings at once and it also acted as a conservation method. However, the population was not finished rising — it rose from 58,000 in 1852 to 267,000 by 1901.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Cistern",
                                        "text": "cisterns"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "Montreal was the capital of the Province of Canada from 1844 to 1849, but lost its status when a Tory mob burnt down the Parliament building to protest the passage of the Rebellion Losses Bill . Thereafter, the capital rotated between Quebec City and Toronto until in 1857, Queen Victoria herself established Ottawa as the capital due to strategic reasons. The reasons were twofold. First, because it was located more in the interior of the Province of Canada, it was less susceptible to attack from the United States. Second, and perhaps more importantly, because it lay on the border between French and English Canada, Ottawa was seen as a compromise between Montreal, Toronto, Kingston and Quebec City, which were all vying to become the young nation's official capital. Ottawa retained the status as capital of Canada when the Province of Canada joined with Nova Scotia and New Brunswick to form the Dominion of Canada in 1867.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Province_of_Canada",
                                        "text": "Province of Canada"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Toryism#Canada",
                                        "text": "Tory"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Burning_of_the_Parliament_Buildings_in_Montreal",
                                        "text": "burnt down the Parliament building"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Rebellion_Losses_Bill",
                                        "text": "Rebellion Losses Bill"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Toronto",
                                        "text": "Toronto"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Queen_Victoria",
                                        "text": "Queen Victoria"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Ottawa",
                                        "text": "Ottawa"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Kingston,_Ontario",
                                        "text": "Kingston"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Nova_Scotia",
                                        "text": "Nova Scotia"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/New_Brunswick",
                                        "text": "New Brunswick"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "An internment camp was set up at Immigration Hall in Montreal from August 1914 to November 1918.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Internment_camp",
                                        "text": "internment camp"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "After World War I, the prohibition movement in the United States led to Montreal becoming a destination for Americans looking for alcohol . Unemployment remained high in the city and was exacerbated by the Stock Market Crash of 1929 and the Great Depression .",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/World_War_I",
                                        "text": "World War I"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Prohibition",
                                        "text": "prohibition"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Alcoholic_beverage",
                                        "text": "alcohol"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Wall_Street_Crash_1929",
                                        "text": "Stock Market Crash of 1929"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Great_Depression",
                                        "text": "Great Depression"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "During World War II, Mayor Camillien Houde protested against conscription and urged Montrealers to disobey the federal government's registry of all men and women. The federal government, part of the Allied forces, was furious over Houde's stand and held him in a prison camp until 1944. That year, the government decided to institute conscription to expand the armed forces and fight the Axis powers . (See Conscription Crisis of 1944 .)",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/World_War_II",
                                        "text": "World War II"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Camillien_Houde",
                                        "text": "Camillien Houde"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Allies_of_World_War_II",
                                        "text": "Allied forces"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Axis_powers",
                                        "text": "Axis powers"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Conscription_Crisis_of_1944",
                                        "text": "Conscription Crisis of 1944"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "Montreal was the official residence of the Luxembourg royal family in exile during World War II.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Luxembourg",
                                        "text": "Luxembourg"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "By 1951, Montreal's population had surpassed one million. However, Toronto's growth had begun challenging Montreal's status as the economic capital of Canada. Indeed, the volume of stocks traded at the Toronto Stock Exchange had already surpassed that traded at the Montreal Stock Exchange in the 1940s. The Saint Lawrence Seaway opened in 1959, allowing vessels to bypass Montreal. In time, this development led to the end of the city's economic dominance as businesses moved to other areas. During the 1960s, there was continued growth as Canada's tallest skyscrapers, new expressways and the subway system known as the Montreal Metro were finished during this time. Montreal also held the World's Fair of 1967, better known as Expo67 .",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Toronto_Stock_Exchange",
                                        "text": "Toronto Stock Exchange"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Montreal_Stock_Exchange",
                                        "text": "Montreal Stock Exchange"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Saint_Lawrence_Seaway",
                                        "text": "Saint Lawrence Seaway"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Montreal_Metro",
                                        "text": "Montreal Metro"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Expo67",
                                        "text": "Expo67"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "The 1970s ushered in a period of wide-ranging social and political changes, stemming largely from the concerns of the French-speaking majority about the conservation of their culture and language, given the traditional predominance of the English Canadian minority in the business arena. The October Crisis and the 1976 election of the Parti Québécois, which supported sovereign status for Quebec, resulted in the departure of many businesses and people from the city. In 1976, Montreal hosted the Summer Olympics . While the event brought the city international prestige and attention, the Olympic Stadium built for the event resulted in massive debt for the city. During the 1980s and early 1990s, Montreal experienced a slower rate of economic growth than many other major Canadian cities. Montreal was the site of the 1989 École Polytechnique massacre, one of Canada's worst mass shootings, where 25-year-old Marc Lépine shot and killed 14 people, all of them women, and wounded 14 other people before shooting himself at École Polytechnique .",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/French_Canadians",
                                        "text": "French-speaking"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/English_Canadians",
                                        "text": "English Canadian"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/October_Crisis",
                                        "text": "October Crisis"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Parti_Québécois",
                                        "text": "Parti Québécois"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/1976_Summer_Olympics",
                                        "text": "Summer Olympics"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Olympic_Stadium_(Montreal)",
                                        "text": "Olympic Stadium"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/École_Polytechnique_massacre",
                                        "text": "École Polytechnique massacre"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Mass_shooting",
                                        "text": "mass shootings"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Marc_Lépine",
                                        "text": "Marc Lépine"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Polytechnique_Montréal",
                                        "text": "École Polytechnique"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "Montreal was merged with the 27 surrounding municipalities on the Island of Montreal on January 1, 2002, creating a unified city encompassing the entire island. There was substantial resistance from the suburbs to the merger, with the perception being that it was forced on the mostly English suburbs by the Parti Québécois. As expected, this move proved unpopular and several mergers were later rescinded. Several former municipalities, totalling 13% of the population of the island, voted to leave the unified city in separate referendums in June 2004. The demerger took place on January 1, 2006, leaving 15 municipalities on the island, including Montreal. Demerged municipalities remain affiliated with the city through an agglomeration council that collects taxes from them to pay for numerous shared services. The 2002 mergers were not the first in the city's history. Montreal annexed 27 other cities, towns and villages beginning with Hochelaga in 1883, with the last prior to 2002 being Pointe-aux-Trembles in 1982.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Montreal_Merger",
                                        "text": "merged"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Referendum",
                                        "text": "referendums"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Mercier–Hochelaga-Maisonneuve#History",
                                        "text": "Hochelaga"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Pointe-aux-Trembles",
                                        "text": "Pointe-aux-Trembles"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "The 21st century has brought with it a revival of the city's economic and cultural landscape. The construction of new residential skyscrapers, two super-hospitals (the Centre hospitalier de l'Université de Montréal and McGill University Health Centre), the creation of the Quartier des Spectacles, reconstruction of the Turcot Interchange, reconfiguration of the Decarie and Dorval interchanges, construction of the new Réseau express métropolitain, gentrification of Griffintown, subway line extensions and the purchase of new subway cars, the complete revitalization and expansion of Trudeau International Airport, the completion of Quebec Autoroute 30, the reconstruction of the Champlain Bridge and the construction of a new toll bridge to Laval are helping Montreal continue to grow.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Centre_hospitalier_de_l'Université_de_Montréal",
                                        "text": "Centre hospitalier de l'Université de Montréal"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/McGill_University_Health_Centre",
                                        "text": "McGill University Health Centre"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Quartier_des_Spectacles",
                                        "text": "Quartier des Spectacles"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Turcot_Interchange",
                                        "text": "Turcot Interchange"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Réseau_express_métropolitain",
                                        "text": "Réseau express métropolitain"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Griffintown",
                                        "text": "Griffintown"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Montréal–Trudeau_International_Airport",
                                        "text": "Trudeau International Airport"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Quebec_Autoroute_30",
                                        "text": "Quebec Autoroute 30"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Champlain_Bridge,_Montreal",
                                        "text": "Champlain Bridge"
                                    }
                                ]
                            }
                        ]
                    }
                ]
            },
            {
                "name": "Geography",
                "type": "section",
                "has_parts": [
                    {
                        "type": "paragraph",
                        "value": "Montreal is in the southwest of the province of Quebec. The city covers most of the Island of Montreal at the confluence of the Saint Lawrence and Ottawa Rivers. The port of Montreal lies at one end of the Saint Lawrence Seaway, the river gateway that stretches from the Great Lakes to the Atlantic. Montreal is defined by its location between the Saint Lawrence river to its south and the Rivière des Prairies to its north. The city is named after the most prominent geographical feature on the island, a three-head hill called Mount Royal, topped at 232 m (761 ft) above sea level .",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Great_Lakes",
                                "text": "Great Lakes"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Rivière_des_Prairies",
                                "text": "Rivière des Prairies"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Above_mean_sea_level",
                                "text": "above sea level"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "Montreal is at the centre of the Montreal Metropolitan Community, and is bordered by the city of Laval to the north; Longueuil, Saint-Lambert, Brossard, and other municipalities to the south; Repentigny to the east and the West Island municipalities to the west. The anglophone enclaves of Westmount, Montreal West, Hampstead, Côte Saint-Luc, the Town of Mount Royal and the francophone enclave Montreal East are all surrounded by Montreal.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Greater_Montreal",
                                "text": "Montreal Metropolitan Community"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Laval,_Quebec",
                                "text": "Laval"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Longueuil",
                                "text": "Longueuil"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Saint-Lambert,_Quebec",
                                "text": "Saint-Lambert"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Brossard",
                                "text": "Brossard"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Repentigny,_Quebec",
                                "text": "Repentigny"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/West_Island",
                                "text": "West Island"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/English_language",
                                "text": "anglophone"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Westmount,_Quebec",
                                "text": "Westmount"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Montreal_West,_Quebec",
                                "text": "Montreal West"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Hampstead,_Quebec",
                                "text": "Hampstead"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Côte_Saint-Luc",
                                "text": "Côte Saint-Luc"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Mount_Royal,_Quebec",
                                "text": "Town of Mount Royal"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Francophone",
                                "text": "francophone"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Montréal-Est,_Quebec",
                                "text": "Montreal East"
                            }
                        ]
                    },
                    {
                        "name": "Climate",
                        "type": "section",
                        "has_parts": [
                            {
                                "type": "paragraph",
                                "value": "Montreal is classified as a warm-summer humid continental climate (Köppen climate classification: Dfb).Summers are warm to hot and humid with a daily maximum average of 26 to 27 °C (79 to 81 °F) in July; temperatures in excess of 30 °C (86 °F) are common. Conversely, cold fronts can bring crisp, drier and windy weather in the early and later parts of summer.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Warm-summer_humid_continental_climate",
                                        "text": "warm-summer humid continental climate"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Köppen_climate_classification",
                                        "text": "Köppen climate classification"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "Winter brings cold, snowy, windy, and, at times, icy weather, with a daily average ranging from −10.5 to −9 °C (13.1 to 15.8 °F) in January. However, some winter days rise above freezing, allowing for rain on an average of 4 days in January and February each. Usually, snow covering some or all bare ground lasts on average from the first or second week of December until the last week of March. While the air temperature does not fall below −30 °C (−22 °F) every year, the wind chill often makes the temperature feel this low to exposed skin.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Wind_chill",
                                        "text": "wind chill"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "Spring and fall are pleasantly mild but prone to drastic temperature changes; spring even more so than fall. Late season heat waves as well as \"Indian summers \" are possible. Early and late season snow storms can occur in November and March, and more rarely in April. Montreal is generally snow free from late April to late October. However, snow can fall in early to mid-October as well as early to mid-May on rare occasions.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Indian_summer",
                                        "text": "Indian summers"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "The lowest temperature in Environment Canada's books was −37.8 °C (−36 °F) on January 15, 1957, and the highest temperature was 37.6 °C (99.7 °F) on August 1, 1975, both at Dorval International Airport .",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Montréal–Pierre_Elliott_Trudeau_International_Airport",
                                        "text": "Dorval International Airport"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "Before modern weather record keeping (which dates back to 1871 for McGill), a minimum temperature almost 5 degrees lower was recorded at 7 a.m. on January 10, 1859, where it registered at −42 °C (−44 °F) ."
                            },
                            {
                                "type": "paragraph",
                                "value": "Annual precipitation is around 1,000 mm (39 in), including an average of about 210 cm (83 in) of snowfall, which occurs from November through March. Thunderstorms are common in the period beginning in late spring through summer to early fall; additionally, tropical storms or their remnants can cause heavy rains and gales. Montreal averages 2,050 hours of sunshine annually, with summer being the sunniest season, though slightly wetter than the others in terms of total precipitation—mostly from thunderstorms."
                            }
                        ]
                    }
                ]
            },
            {
                "name": "Architecture",
                "type": "section",
                "has_parts": [
                    {
                        "type": "paragraph",
                        "value": "For over a century and a half, Montreal was the industrial and financial centre of Canada. This legacy has left a variety of buildings including factories, elevators, warehouses, mills, and refineries, that today provide an invaluable insight into the city's history, especially in the downtown area and the Old Port area. There are 50 National Historic Sites of Canada, more than any other city.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Grain_elevators",
                                "text": "elevators"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Warehouses",
                                "text": "warehouses"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Refineries",
                                "text": "refineries"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/List_of_National_Historic_Sites_of_Canada_in_Montreal",
                                "text": "National Historic Sites of Canada"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "Some of the city's earliest still-standing buildings date back to the late 17th and early 18th centuries. Although most are clustered around the Old Montreal area, such as the Sulpician Seminary adjacent to Notre Dame Basilica that dates back to 1687, and Château Ramezay, which was built in 1705, examples of early colonial architecture are dotted throughout the city. Situated in Lachine, the Le Ber-Le Moyne House is the oldest complete building in the city, built between 1669 and 1671. In Point St. Charles visitors can see the Maison Saint-Gabriel, which can trace its history back to 1698. There are many historic buildings in Old Montreal in their original form: Notre Dame of Montreal Basilica, Bonsecours Market, and the 19th‑century headquarters of all major Canadian banks on St. James Street (French: Rue Saint Jacques). Montreal's earliest buildings are characterized by their uniquely French influence and grey stone construction.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Notre-Dame_de_Montréal_Basilica",
                                "text": "Notre Dame of Montreal Basilica"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Bonsecours_Market",
                                "text": "Bonsecours Market"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Saint_Jacques_Street",
                                "text": "St. James Street"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "Saint Joseph's Oratory, completed in 1967, Ernest Cormier 's Art Deco Université de Montréal main building, the landmark Place Ville Marie office tower, the controversial Olympic Stadium and surrounding structures, are but a few notable examples of the city's 20th-century architecture. Pavilions designed for the 1967 International and Universal Exposition, popularly known as Expo 67, featured a wide range of architectural designs. Though most pavilions were temporary structures, several have become landmarks, including Buckminster Fuller's geodesic dome U.S. Pavilion, now the Montreal Biosphere, and Moshe Safdie 's striking Habitat 67 apartment complex.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Saint_Joseph's_Oratory",
                                "text": "Saint Joseph's Oratory"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Ernest_Cormier",
                                "text": "Ernest Cormier"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Art_Deco",
                                "text": "Art Deco"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Université_de_Montréal",
                                "text": "Université de Montréal"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Place_Ville_Marie",
                                "text": "Place Ville Marie"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Geodesic_dome",
                                "text": "geodesic dome"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Montreal_Biosphere",
                                "text": "Montreal Biosphere"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Moshe_Safdie",
                                "text": "Moshe Safdie"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "The Montreal Metro has public artwork by some of the biggest names in Quebec culture .",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Culture_of_Quebec",
                                "text": "Quebec culture"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "In 2006 Montreal was named a UNESCO City of Design, one of only three design capitals of the world (the others being Berlin and Buenos Aires). This distinguished title recognizes Montreal's design community. Since 2005 the city has been home for the International Council of Graphic Design Associations (Icograda); the International Design Alliance (IDA).",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Berlin",
                                "text": "Berlin"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Buenos_Aires",
                                "text": "Buenos Aires"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/International_Council_of_Graphic_Design_Associations",
                                "text": "International Council of Graphic Design Associations"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "The Underground City (officially RÉSO), an important tourist attraction, is an underground network connecting shopping centres, pedestrian thoroughfares, universities, hotels, restaurants, bistros, subway stations and more, in and around downtown with 32 km (20 mi) of tunnels over 12 km (4.6 sq mi) in the most densely populated part of Montreal.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Underground_City,_Montreal",
                                "text": "Underground City"
                            }
                        ]
                    }
                ]
            },
            {
                "name": "Neighbourhoods",
                "type": "section",
                "has_parts": [
                    {
                        "type": "paragraph",
                        "value": "The city is composed of 19 large boroughs, subdivided into neighbourhoods.The boroughs are:Côte-des-Neiges–Notre-Dame-de-Grace, The Plateau Mount Royal, Outremont and Ville Marie in the centre; Mercier–Hochelaga-Maisonneuve, Rosemont–La Petite-Patrie and Villeray–Saint-Michel–Parc-Extension in the east; Anjou, Montréal-Nord, Rivière-des-Prairies–Pointe-aux-Trembles and Saint-Leonard in the northeast; Ahuntsic-Cartierville, L'Île-Bizard–Sainte-Geneviève, Pierrefonds-Roxboro and Saint-Laurent in the northwest; and Lachine, LaSalle, The South West and Verdun in the south.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Boroughs_of_Montreal",
                                "text": "boroughs"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Côte-des-Neiges–Notre-Dame-de-Grâce",
                                "text": "Côte-des-Neiges–Notre-Dame-de-Grace"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Le_Plateau-Mont-Royal",
                                "text": "The Plateau Mount Royal"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Outremont_(borough)",
                                "text": "Outremont"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Ville-Marie_(Montreal)",
                                "text": "Ville Marie"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Mercier–Hochelaga-Maisonneuve",
                                "text": "Mercier–Hochelaga-Maisonneuve"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Rosemont–La_Petite-Patrie",
                                "text": "Rosemont–La Petite-Patrie"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Villeray–Saint-Michel–Parc-Extension",
                                "text": "Villeray–Saint-Michel–Parc-Extension"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Anjou_(borough)",
                                "text": "Anjou"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Montréal-Nord",
                                "text": "Montréal-Nord"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Rivière-des-Prairies–Pointe-aux-Trembles",
                                "text": "Rivière-des-Prairies–Pointe-aux-Trembles"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Saint-Leonard,_Quebec",
                                "text": "Saint-Leonard"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Ahuntsic-Cartierville",
                                "text": "Ahuntsic-Cartierville"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/L'Île-Bizard–Sainte-Geneviève",
                                "text": "L'Île-Bizard–Sainte-Geneviève"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Pierrefonds-Roxboro",
                                "text": "Pierrefonds-Roxboro"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Saint-Laurent_(borough)",
                                "text": "Saint-Laurent"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Lachine_(borough)",
                                "text": "Lachine"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/LaSalle_(borough)",
                                "text": "LaSalle"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Le_Sud-Ouest",
                                "text": "The South West"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Verdun_(borough)",
                                "text": "Verdun"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "Many of these boroughs were independent cities that were forced to merge with Montreal in January 2002 following the 2002 municipal reorganization of Montreal .",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/2002-2006_municipal_reorganization_of_Montreal",
                                "text": "2002 municipal reorganization of Montreal"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "The borough with the most neighbourhoods is Ville Marie, which includes downtown, the historical district of Old Montreal, Chinatown, the Gay Village, the Latin Quarter, the gentrified Quartier international and Cité Multimédia as well as the Quartier des Spectacles which is under development. Other neighbourhoods of interest in the borough include the affluent Golden Square Mile neighbourhood at the foot of Mount Royal and the Shaughnessy Village /Concordia U area home to thousands of students at Concordia University . The borough also comprises most of Mount Royal Park, Saint Helen's Island, and Notre-Dame Island .",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Chinatown_(Montreal)",
                                "text": "Chinatown"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Gay_Village,_Montreal",
                                "text": "Gay Village"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Quartier_Latin,_Montreal",
                                "text": "Latin Quarter"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Quartier_international_de_Montréal",
                                "text": "Quartier international"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Cité_Multimédia",
                                "text": "Cité Multimédia"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Shaughnessy_Village",
                                "text": "Shaughnessy Village"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Quartier_Concordia",
                                "text": "Concordia U"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Concordia_University_(Montreal)",
                                "text": "Concordia University"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Mount_Royal_Park",
                                "text": "Mount Royal Park"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Saint_Helen's_Island",
                                "text": "Saint Helen's Island"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Île_Notre-Dame",
                                "text": "Notre-Dame Island"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "The Plateau Mount Royal borough was a working class francophone area. The largest neighbourhood is the Plateau (not to be confused with the whole borough), which is undergoing considerable gentrification, and a 2001 study deemed it as Canada's most creative neighbourhood because artists comprise 8% of its labour force. The neighbourhood of Mile End in the northwestern part of the borough has been a very multicultural area of the city, and features two of Montreal's well-known bagel establishments, St-Viateur Bagel and Fairmount Bagel . The McGill Ghetto is in the extreme southwestern portion of the borough, its name being derived from the fact that it is home to thousands of McGill University students and faculty members.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Plateau",
                                "text": "Plateau"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/The_Plateau",
                                "text": "the Plateau"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Mile_End,_Montreal",
                                "text": "Mile End"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Montreal-style_bagel",
                                "text": "bagel establishments"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/St-Viateur_Bagel",
                                "text": "St-Viateur Bagel"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Fairmount_Bagel",
                                "text": "Fairmount Bagel"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/McGill_Ghetto",
                                "text": "McGill Ghetto"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/McGill_University",
                                "text": "McGill University"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "The South West borough was home to much of the city's industry during the late 19th and early-to-mid 20th century. The borough included Goose Village and was historically home to the traditionally working-class Irish neighbourhoods of Griffintown and Point Saint Charles as well as the low-income neighbourhoods of Saint Henri and Little Burgundy .",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Goose_Village,_Montreal",
                                "text": "Goose Village"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Irish_Quebecer",
                                "text": "Irish"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Pointe-Saint-Charles",
                                "text": "Point Saint Charles"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Saint-Henri",
                                "text": "Saint Henri"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Little_Burgundy",
                                "text": "Little Burgundy"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "Other notable neighbourhoods include the multicultural areas of Notre-Dame-de-Grâce and Côte-des-Neiges in the Côte-des-Neiges–Notre-Dame-de-Grace borough, and Little Italy in the borough of Rosemont–La Petite-Patrie and Hochelaga-Maisonneuve, home of the Olympic Stadium in the borough of Mercier–Hochelaga-Maisonneuve.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Notre-Dame-de-Grâce",
                                "text": "Notre-Dame-de-Grâce"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Côte-des-Neiges",
                                "text": "Côte-des-Neiges"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Little_Italy,_Montreal",
                                "text": "Little Italy"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Hochelaga-Maisonneuve",
                                "text": "Hochelaga-Maisonneuve"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Olympic_Stadium,_Montreal",
                                "text": "Olympic Stadium"
                            }
                        ]
                    },
                    {
                        "name": "Old Montreal",
                        "type": "section",
                        "has_parts": [
                            {
                                "type": "paragraph",
                                "value": "Old Montreal is a historic area southeast of downtown containing many attractions such as the Old Port of Montreal, Place Jacques-Cartier, Montreal City Hall, the Bonsecours Market, Place d'Armes, Pointe-à-Callière Museum, the Notre-Dame de Montréal Basilica, and the Montreal Science Centre .",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Old_Port_of_Montreal",
                                        "text": "Old Port of Montreal"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Montreal_City_Hall",
                                        "text": "Montreal City Hall"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Bonsecours_Market",
                                        "text": "Bonsecours Market"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Pointe-à-Callière_Museum",
                                        "text": "Pointe-à-Callière Museum"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Montreal_Science_Centre",
                                        "text": "Montreal Science Centre"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "Architecture and cobbled streets in Old Montreal have been maintained or restored. Old Montreal is accessible from the downtown core via the underground city and is served by several STM bus routes and Metro stations, ferries to the South Shore and a network of bicycle paths.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Underground_city,_Montreal",
                                        "text": "underground city"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Société_de_transport_de_Montréal",
                                        "text": "STM"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "The riverside area adjacent to Old Montreal is known as the Old Port. The Old Port was the site of the Port of Montreal, but its shipping operations have been moved to a larger site downstream, leaving the former location as a recreational and historical area maintained by Parks Canada . The new Port of Montreal is Canada's largest container port and the largest inland port on Earth.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Port_of_Montreal",
                                        "text": "Port of Montreal"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Parks_Canada",
                                        "text": "Parks Canada"
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "name": "Mount Royal",
                        "type": "section",
                        "has_parts": [
                            {
                                "type": "paragraph",
                                "value": "The mountain is the site of Mount Royal Park, one of Montreal's largest greenspaces . The park, most of which is wooded, was designed by Frederick Law Olmsted, who also designed New York's Central Park, and was inaugurated in 1876.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Open_space_reserve",
                                        "text": "greenspaces"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Frederick_Law_Olmsted",
                                        "text": "Frederick Law Olmsted"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Central_Park",
                                        "text": "Central Park"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "The park contains two belvederes, the more prominent of which is the Kondiaronk Belvedere, a semicircular plaza with a chalet overlooking Downtown Montreal. Other features of the park are Beaver Lake, a small man-made lake, a short ski slope, a sculpture garden, Smith House, an interpretive centre, and a well-known monument to Sir George-Étienne Cartier . The park hosts athletic, tourist and cultural activities.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Belvedere_(structure)",
                                        "text": "belvederes"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Chalet",
                                        "text": "chalet"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Skiing",
                                        "text": "ski"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Ski_slope",
                                        "text": "slope"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Sculpture_garden",
                                        "text": "sculpture garden"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Interpretive_centre",
                                        "text": "interpretive centre"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/George-Étienne_Cartier_Monument",
                                        "text": "monument to Sir George-Étienne Cartier"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "The mountain is home to two major cemeteries, Notre-Dame-des-Neiges (founded in 1854) and Mount Royal (1852). Mount Royal Cemetery is a 165 acres (67 ha) terraced cemetery on the north slope of Mount Royal in the borough of Outremont. Notre Dame des Neiges Cemetery is much larger, predominantly French-Canadian and officially Catholic. More than 900,000 people are buried there.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Mount_Royal_Cemetery",
                                        "text": "Mount Royal Cemetery"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Notre_Dame_des_Neiges_Cemetery",
                                        "text": "Notre Dame des Neiges Cemetery"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "Mount Royal Cemetery contains more than 162,000 graves and is the final resting place for a number of notable Canadians. It includes a veterans section with several soldiers who were awarded the British Empire 's highest military honour, the Victoria Cross . In 1901, the Mount Royal Cemetery Company established the first crematorium in Canada.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/British_Empire",
                                        "text": "British Empire"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Victoria_Cross",
                                        "text": "Victoria Cross"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "The first cross on the mountain was placed there in 1643 by Paul Chomedey de Maisonneuve, the founder of the city, in fulfilment of a vow he made to the Virgin Mary when praying to her to stop a disastrous flood. Today, the mountain is crowned by a 31.4 m-high (103 ft) illuminated cross, installed in 1924 by the John the Baptist Society and now owned by the city. It was converted to fibre optic light in 1992. The new system can turn the lights red, blue, or purple, the last of which is used as a sign of mourning between the death of the Pope and the election of the next.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Mount_Royal_Cross",
                                        "text": "cross"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Mary,_mother_of_Jesus",
                                        "text": "Virgin Mary"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Prayer",
                                        "text": "praying"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Société_Saint-Jean-Baptiste",
                                        "text": "John the Baptist Society"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Optical_fibre",
                                        "text": "fibre optic"
                                    }
                                ]
                            }
                        ]
                    }
                ]
            },
            {
                "name": "Demographics",
                "type": "section",
                "has_parts": [
                    {
                        "type": "paragraph",
                        "value": "In the 2021 Census of Population conducted by Statistics Canada, Montréal had a population of 1,762,949 living in 816,338 of its 878,542 total private dwellings, a change of 3.4% from its 2016 population of 1,704,694 . With a land area of 364.74 km (140.83 sq mi), it had a population density of 4,833.4/km (12,518.6/sq mi) in 2021.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/2021_Canadian_census",
                                "text": "2021 Census of Population"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Statistics_Canada",
                                "text": "Statistics Canada"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "According to Statistics Canada, at the 2016 Canadian census the city had 1,704,694 inhabitants. A total of 4,098,927 lived in the Montreal Census Metropolitan Area (CMA) at the same 2016 census, up from 3,934,078 at the 2011 census (within 2011 CMA boundaries), which is a population growth of 4.19% from 2011 to 2016. In 2015, the Greater Montreal population was estimated at 4,060,700. According to StatsCan, by 2030, the Greater Montreal Area is expected to number 5,275,000 with 1,722,000 being visible minorities.In the 2016 census, children under 14 years of age (691,345) constituted 16.9%, while inhabitants over 65 years of age (671,690) numbered 16.4% of the total population of the CMA.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Statistics_Canada",
                                "text": "Statistics Canada"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/2016_Canadian_census",
                                "text": "2016 Canadian census"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Census_Metropolitan_Area",
                                "text": "Census Metropolitan Area"
                            }
                        ]
                    },
                    {
                        "name": "Ethnicity",
                        "type": "section",
                        "has_parts": [
                            {
                                "type": "paragraph",
                                "value": "People of European ethnicities formed the largest cluster of ethnic groups. The largest reported European ethnicities in the 2006 census were French 23%, Italians 10%, Irish 5%, English 4%, Scottish 3%, and Spanish 2%.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/European_ethnic_groups",
                                        "text": "European ethnicities"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/French_people",
                                        "text": "French"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Italian_people",
                                        "text": "Italians"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Irish_people",
                                        "text": "Irish"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/English_people",
                                        "text": "English"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Scottish_people",
                                        "text": "Scottish"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Spanish_people",
                                        "text": "Spanish"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "The panethnic breakdown of the city of Montreal as per the 2021 census was European (1,038,940 residents or 60.3% of the population), African (198,610; 11.5%), Middle Eastern (159,435; 9.3%), South Asian (79,670; 4.6%), Latin American (78,150; 4.5%), Southeast Asian (65,260; 3.8%), East Asian (64,825; 3.8%), Indigenous (15,315; 0.9%), and Other/Multiracial (23,010; 1.3%).",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Panethnicity",
                                        "text": "panethnic"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/2021_Canadian_census",
                                        "text": "2021 census"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/European_Canadians",
                                        "text": "European"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/African-Canadian",
                                        "text": "African"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Middle_Eastern_Canadians",
                                        "text": "Middle Eastern"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/South_Asian_Canadians",
                                        "text": "South Asian"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Latin_American_Canadians",
                                        "text": "Latin American"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Southeast_Asia",
                                        "text": "Southeast Asian"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/East_Asian_Canadians",
                                        "text": "East Asian"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Indigenous_Canadian",
                                        "text": "Indigenous"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Multiracial_people",
                                        "text": "Multiracial"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "Visible minorities comprised 38.8% of the city of Montreal population in the 2021 census . The five most numerous visible minorities are Black Canadians (11.5%), Arab Canadians (8.2%), South Asian Canadians (4.6%), Latin Americans (4.5%), and Chinese Canadians (3.3%). Furthermore, some 27.2% of the population Greater Montreal are members of a visible minority group as of 2021, up from 5.2% in 1981. Visible minorities are defined by the Canadian Employment Equity Act as \"persons, other than Aboriginals, who are non-white in colour\".",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Visible_minorities",
                                        "text": "Visible minorities"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Canada_2021_Census",
                                        "text": "2021 census"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Black_Canadians",
                                        "text": "Black Canadians"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Arab_Canadians",
                                        "text": "Arab Canadians"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Chinese_Canadians",
                                        "text": "Chinese Canadians"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Greater_Montreal",
                                        "text": "Greater Montreal"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Canadian_Employment_Equity_Act",
                                        "text": "Canadian Employment Equity Act"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Aboriginal_peoples_of_Canada",
                                        "text": "Aboriginals"
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "name": "Language",
                        "type": "section",
                        "has_parts": [
                            {
                                "type": "paragraph",
                                "value": "As of the 2021 Census, 47.0% of Montreal residents spoke French alone as a first language, while 13.0% spoke English alone. 2% spoke both English and French as first languages, 2.6% spoke both French and a non-official language and 1.5% spoke both English and a non-official language. 0.8% of residents spoke English, French and a non-official language as first languages. 32.8% of residents had a non-official language as a mother tongue. The most common were Arabic (5.7%), Spanish (4.6%), Italian (3.3%), Chinese Languages (2.7%), Haitian Creole (1.6%), Vietnamese (1.1%), and Portuguese (1.0%).",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/French_language",
                                        "text": "French"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/English_language",
                                        "text": "English"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Arabic",
                                        "text": "Arabic"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Spanish_language",
                                        "text": "Spanish"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Italian_language",
                                        "text": "Italian"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Chinese_language",
                                        "text": "Chinese Languages"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Haitian_Creole",
                                        "text": "Haitian Creole"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Vietnamese_language",
                                        "text": "Vietnamese"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Portuguese_language",
                                        "text": "Portuguese"
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "name": "Immigration",
                        "type": "section",
                        "has_parts": [
                            {
                                "type": "paragraph",
                                "value": "The 2021 census reported that immigrants (individuals born outside Canada) comprise 576,125 persons or 33.4% of the total population of Montreal. Of the total immigrant population, the top countries of origin were Haiti (47,550 residents or 8.3% of the population), Algeria (43,840; 7.6%), France (39,275; 6.8%), Morocco (33,005; 5.7%), Italy (30,215; 5.2%), China (26,335; 4.6%), the Philippines (20,475; 3.6%), Lebanon (17,455; 3.0%), Vietnam (16,395; 2.8%), and India (13,575; 2.4%).",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/2021_Canadian_census",
                                        "text": "2021 census"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Immigration_to_Canada",
                                        "text": "immigrants"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Haiti",
                                        "text": "Haiti"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Algeria",
                                        "text": "Algeria"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/France",
                                        "text": "France"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Morocco",
                                        "text": "Morocco"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Italy",
                                        "text": "Italy"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/China",
                                        "text": "China"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Philippines",
                                        "text": "Philippines"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Lebanon",
                                        "text": "Lebanon"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Vietnam",
                                        "text": "Vietnam"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/India",
                                        "text": "India"
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "name": "Religion",
                        "type": "section",
                        "has_parts": [
                            {
                                "type": "paragraph",
                                "value": "The Greater Montreal Area is predominantly Catholic; however, weekly church attendance in Quebec was among the lowest in Canada in 1998. Historically Montreal has been a centre of Catholicism in North America with its numerous seminaries and churches, including the Notre-Dame Basilica, the Cathédrale Marie-Reine-du-Monde, and Saint Joseph's Oratory .",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Catholic_Church",
                                        "text": "Catholic"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Notre-Dame_Basilica_(Montreal)",
                                        "text": "Notre-Dame Basilica"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Cathédrale_Marie-Reine-du-Monde",
                                        "text": "Cathédrale Marie-Reine-du-Monde"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Saint_Joseph's_Oratory",
                                        "text": "Saint Joseph's Oratory"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "Some 49.5% of the total population is Christian, largely Roman Catholic (35.0%), primarily because of descendants of original French settlers, and others of Italian and Irish origins. Protestants which include Anglican Church in Canada, United Church of Canada, Lutheran, owing to British and German immigration, and other denominations number 11.3%, with a further 3.2% consisting mostly of Orthodox Christians, fuelled by a large Greek population. There is also a number of Russian and Ukrainian Orthodox parishes.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Protestants",
                                        "text": "Protestants"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Anglican_Church_in_Canada",
                                        "text": "Anglican Church in Canada"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/United_Church_of_Canada",
                                        "text": "United Church of Canada"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Lutheran",
                                        "text": "Lutheran"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Eastern_Orthodox",
                                        "text": "Orthodox Christians"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "Islam is the largest non-Christian religious group, with 218,395 members, the second-largest concentration of Muslims in Canada at 12.7%. The Jewish community in Montreal has a population of 90,780. In cities such as Côte Saint-Luc and Hampstead, Jewish people constitute the majority, or a substantial part of the population. In 1971 the Jewish community in Greater Montreal numbered 109,480. Political and economic uncertainties led many to leave Montreal and the province of Quebec.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Islam",
                                        "text": "Islam"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/History_of_the_Jews_in_Montreal",
                                        "text": "Jewish community in Montreal"
                                    }
                                ]
                            }
                        ]
                    }
                ]
            },
            {
                "name": "Economy",
                "type": "section",
                "has_parts": [
                    {
                        "type": "paragraph",
                        "value": "Montreal has the second-largest economy of Canadian cities based on GDP and the largest in Quebec. In 2019, Metropolitan Montreal was responsible for CA$234.0 billion of Quebec's CA$425.3 billion GDP. The city is today an important centre of commerce, finance, industry, technology, culture, world affairs and is the headquarters of the Montreal Exchange . In recent decades, the city was widely seen as weaker than that of Toronto and other major Canadian cities, but it has recently experienced a revival.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Montreal_Exchange",
                                "text": "Montreal Exchange"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "Industries include aerospace, electronic goods, pharmaceuticals, printed goods, software engineering, telecommunications, textile and apparel manufacturing, tobacco, petrochemicals, and transportation. The service sector is also strong and includes civil, mechanical and process engineering, finance, higher education, and research and development. In 2002, Montreal was the fourth-largest centre in North America in terms of aerospace jobs.The Port of Montreal is one of the largest inland ports in the world handling 26 million tonnes of cargo annually. As one of the most important ports in Canada, it remains a transshipment point for grain, sugar, petroleum products, machinery, and consumer goods. For this reason, Montreal is the railway hub of Canada and has always been an extremely important rail city; it is home to the headquarters of the Canadian National Railway, and was home to the headquarters of the Canadian Pacific Railway until 1995.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Aerospace",
                                "text": "aerospace"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Electronics",
                                "text": "electronic"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Pharmaceuticals",
                                "text": "pharmaceuticals"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Software_engineering",
                                "text": "software engineering"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Civil_engineering",
                                "text": "civil"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Mechanical_engineering",
                                "text": "mechanical"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Process_engineering",
                                "text": "process engineering"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Cereal",
                                "text": "grain"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Canadian_National_Railway",
                                "text": "Canadian National Railway"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Canadian_Pacific_Railway",
                                "text": "Canadian Pacific Railway"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "The headquarters of the Canadian Space Agency is in Longueuil, southeast of Montreal. Montreal also hosts the headquarters of the International Civil Aviation Organization (ICAO, a United Nations body); the World Anti-Doping Agency (an Olympic body); the Airports Council International (the association of the world's airports – ACI World); the International Air Transport Association (IATA), IATA Operational Safety Audit and the International Gay and Lesbian Chamber of Commerce (IGLCC), as well as some other international organizations in various fields.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Canadian_Space_Agency",
                                "text": "Canadian Space Agency"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/International_Civil_Aviation_Organization",
                                "text": "International Civil Aviation Organization"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/United_Nations",
                                "text": "United Nations"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/World_Anti-Doping_Agency",
                                "text": "World Anti-Doping Agency"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/International_Olympic_Committee",
                                "text": "Olympic"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Airports_Council_International",
                                "text": "Airports Council International"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/International_Air_Transport_Association",
                                "text": "International Air Transport Association"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/IATA_Operational_Safety_Audit",
                                "text": "IATA Operational Safety Audit"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "Montreal is a centre of film and television production. The headquarters of Alliance Films and five studios of the Academy Award -winning National Film Board of Canada are in the city, as well as the head offices of Telefilm Canada, the national feature-length film and television funding agency and Télévision de Radio-Canada . Given its eclectic architecture and broad availability of film services and crew members, Montreal is a popular filming location for feature-length films, and sometimes stands in for European locations. The city is also home to many recognized cultural, film, and music festivals (Just For Laughs, Just For Laughs Gags, Montreal International Jazz Festival, and others), which contribute significantly to its economy. It is also home to one of the world's largest cultural enterprises, the Cirque du Soleil .",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Alliance_Films",
                                "text": "Alliance Films"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Academy_Awards",
                                "text": "Academy Award"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/National_Film_Board_of_Canada",
                                "text": "National Film Board of Canada"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Telefilm_Canada",
                                "text": "Telefilm Canada"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Télévision_de_Radio-Canada",
                                "text": "Télévision de Radio-Canada"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Just_For_Laughs",
                                "text": "Just For Laughs"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Just_For_Laughs_Gags",
                                "text": "Just For Laughs Gags"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Montreal_International_Jazz_Festival",
                                "text": "Montreal International Jazz Festival"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Cirque_du_Soleil",
                                "text": "Cirque du Soleil"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "Montreal is also a global hub for artificial intelligence research with many companies involved in this sector, such as Facebook AI Research (FAIR), Microsoft Research, Google Brain, DeepMind, Samsung Research and Thales Group (cortAIx). The city is also home to Mila (research institute), an artificial intelligence research institute with over 500 researchers specializing in the field of deep learning, the largest of its kind in the world.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Artificial_intelligence",
                                "text": "artificial intelligence"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Facebook",
                                "text": "Facebook"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Microsoft_Research",
                                "text": "Microsoft Research"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Google_Brain",
                                "text": "Google Brain"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/DeepMind",
                                "text": "DeepMind"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Samsung_Research",
                                "text": "Samsung Research"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Thales_Group",
                                "text": "Thales Group"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Mila_(research_institute)",
                                "text": "Mila (research institute)"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "The video game industry has been booming in Montreal since November 2, 1995, coinciding with the opening of Ubisoft Montreal . Recently, the city has attracted world leading game developers and publishers studios such as EA, Eidos Interactive, BioWare, Artificial Mind and Movement, Strategy First, THQ, Gameloft mainly because of the quality of local specialized labour, and tax credits offered to the corporations. In 2010, Warner Bros. Interactive Entertainment, a division of Warner Bros., announced that it would open a video game studio. Relatively new to the video game industry, it will be Warner Bros. first studio opened, not purchased, and will develop games for such Warner Bros. franchises as Batman and other games from their DC Comics portfolio. The studio will create 300 jobs.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Ubisoft_Montreal",
                                "text": "Ubisoft Montreal"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/EA_Montreal",
                                "text": "EA"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Eidos_Interactive",
                                "text": "Eidos Interactive"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/BioWare",
                                "text": "BioWare"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Artificial_Mind_and_Movement",
                                "text": "Artificial Mind and Movement"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Strategy_First",
                                "text": "Strategy First"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/THQ",
                                "text": "THQ"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Gameloft",
                                "text": "Gameloft"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Warner_Bros._Interactive_Entertainment",
                                "text": "Warner Bros. Interactive Entertainment"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Warner_Bros.",
                                "text": "Warner Bros."
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Batman",
                                "text": "Batman"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/DC_Comics",
                                "text": "DC Comics"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "Montreal plays an important role in the finance industry. The sector employs approximately 100,000 people in the Greater Montreal Area. As of March 2018, Montreal is ranked in the 12th position in the Global Financial Centres Index, a ranking of the competitiveness of financial centres around the world. The city is home to the Montreal Exchange, the oldest stock exchange in Canada and the only financial derivatives exchange in the country. The corporate headquarters of the Bank of Montreal and Royal Bank of Canada, two of the biggest banks in Canada, were in Montreal. While both banks moved their headquarters to Toronto, Ontario, their legal corporate offices remain in Montreal. The city is home to head offices of two smaller banks, National Bank of Canada and Laurentian Bank of Canada . The Caisse de dépôt et placement du Québec, an institutional investor managing assets totalling $408 billion CAD, has its main business office in Montreal. Many foreign subsidiaries operating in the financial sector also have offices in Montreal, including HSBC, Aon, Société Générale, BNP Paribas and AXA .",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Global_Financial_Centres_Index",
                                "text": "Global Financial Centres Index"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Ranking",
                                "text": "ranking"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Competition_(companies)",
                                "text": "competitiveness"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Financial_centre",
                                "text": "financial centres"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Bank_of_Montreal",
                                "text": "Bank of Montreal"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Royal_Bank_of_Canada",
                                "text": "Royal Bank of Canada"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/National_Bank_of_Canada",
                                "text": "National Bank of Canada"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Laurentian_Bank_of_Canada",
                                "text": "Laurentian Bank of Canada"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Caisse_de_dépôt_et_placement_du_Québec",
                                "text": "Caisse de dépôt et placement du Québec"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/HSBC",
                                "text": "HSBC"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Aon_(company)",
                                "text": "Aon"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Société_Générale",
                                "text": "Société Générale"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/BNP_Paribas",
                                "text": "BNP Paribas"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/AXA",
                                "text": "AXA"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "Several companies are headquartered in Greater Montreal Area including Rio Tinto Alcan, Bombardier Inc., Canadian National Railway, CGI Group, Air Canada, Air Transat, CAE, Saputo, Cirque du Soleil, Stingray Group, Quebecor, Ultramar, Kruger Inc., Jean Coutu Group, Uniprix, Proxim, Domtar, Le Château, Power Corporation, Cellcom Communications, Bell Canada . Standard Life, Hydro-Québec, AbitibiBowater, Pratt and Whitney Canada, Molson, Tembec, Canada Steamship Lines, Fednav, Alimentation Couche-Tard, SNC-Lavalin, MEGA Brands, Aeroplan, Agropur, Metro Inc., Laurentian Bank of Canada, National Bank of Canada, Transat A.T., Via Rail, GardaWorld, Novacam Technologies, SOLABS, Dollarama, Rona and the Caisse de dépôt et placement du Québec .",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Rio_Tinto_Alcan",
                                "text": "Rio Tinto Alcan"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Bombardier_Inc.",
                                "text": "Bombardier Inc."
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Canadian_National_Railway",
                                "text": "Canadian National Railway"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/CGI_Group",
                                "text": "CGI Group"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Air_Canada",
                                "text": "Air Canada"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Air_Transat",
                                "text": "Air Transat"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/CAE_(company)",
                                "text": "CAE"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Saputo_Incorporated",
                                "text": "Saputo"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Cirque_du_Soleil",
                                "text": "Cirque du Soleil"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Stingray_Group",
                                "text": "Stingray Group"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Quebecor",
                                "text": "Quebecor"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Ultramar",
                                "text": "Ultramar"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Kruger_Inc.",
                                "text": "Kruger Inc."
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Jean_Coutu_Group",
                                "text": "Jean Coutu Group"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Uniprix",
                                "text": "Uniprix"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Proxim_(pharmacy)",
                                "text": "Proxim"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Domtar",
                                "text": "Domtar"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Le_Château",
                                "text": "Le Château"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Power_Corporation",
                                "text": "Power Corporation"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Cellcom_Communications",
                                "text": "Cellcom Communications"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Bell_Canada",
                                "text": "Bell Canada"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Standard_Life_(Canada)",
                                "text": "Standard Life"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Hydro-Québec",
                                "text": "Hydro-Québec"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/AbitibiBowater",
                                "text": "AbitibiBowater"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Pratt_and_Whitney_Canada",
                                "text": "Pratt and Whitney Canada"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Molson",
                                "text": "Molson"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Tembec",
                                "text": "Tembec"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Canada_Steamship_Lines",
                                "text": "Canada Steamship Lines"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Fednav",
                                "text": "Fednav"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Alimentation_Couche-Tard",
                                "text": "Alimentation Couche-Tard"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/SNC-Lavalin",
                                "text": "SNC-Lavalin"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/MEGA_Brands",
                                "text": "MEGA Brands"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Aeroplan",
                                "text": "Aeroplan"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Agropur",
                                "text": "Agropur"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Metro_Inc.",
                                "text": "Metro Inc."
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Laurentian_Bank_of_Canada",
                                "text": "Laurentian Bank of Canada"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/National_Bank_of_Canada",
                                "text": "National Bank of Canada"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Transat_A.T.",
                                "text": "Transat A.T."
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Via_Rail",
                                "text": "Via Rail"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/GardaWorld",
                                "text": "GardaWorld"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Novacam_Technologies",
                                "text": "Novacam Technologies"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Dollarama",
                                "text": "Dollarama"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Rona_(company)",
                                "text": "Rona"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Caisse_de_dépôt_et_placement_du_Québec",
                                "text": "Caisse de dépôt et placement du Québec"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "The Montreal Oil Refining Centre is the largest refining centre in Canada, with companies like Petro-Canada, Ultramar, Gulf Oil, Petromont, Ashland Canada, Parachem Petrochemical, Coastal Petrochemical, Interquisa (Cepsa) Petrochemical, Nova Chemicals, and more. Shell decided to close the refining centre in 2010, throwing hundreds out of work and causing an increased dependence on foreign refineries for eastern Canada.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Montreal_Oil_Refining_Centre",
                                "text": "Montreal Oil Refining Centre"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Petro-Canada",
                                "text": "Petro-Canada"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Ultramar",
                                "text": "Ultramar"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Gulf_Oil",
                                "text": "Gulf Oil"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Cepsa",
                                "text": "Cepsa"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Nova_Chemicals",
                                "text": "Nova Chemicals"
                            }
                        ]
                    }
                ]
            },
            {
                "name": "Culture",
                "type": "section",
                "has_parts": [
                    {
                        "type": "paragraph",
                        "value": "Montreal was referred to as \"Canada's Cultural Capital\" by Monocle magazine . The city is Canada's centre for French-language television productions, radio, theatre, film, multimedia, and print publishing. Montreal's many cultural communities have given it a distinct local culture. Montreal was designated as the World Book Capital for the year 2005 by UNESCO .",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Monocle_(2007_magazine)",
                                "text": "Monocle magazine"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/World_Book_Capital",
                                "text": "World Book Capital"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/UNESCO",
                                "text": "UNESCO"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "Being at the confluence of the French and the English traditions, Montreal has developed a unique and distinguished cultural face. The city has produced much talent in the fields of visual arts, theatre, dance, and music, with a tradition of producing both jazz and rock music. Another distinctive characteristic of cultural life is the vibrancy of its downtown, particularly during summer, prompted by cultural and social events, including its more than 100 annual festivals, the largest being the Montreal International Jazz Festival which is the largest jazz festival in the world. Other popular events include the Just for Laughs (largest comedy festival in the world), Montreal World Film Festival, Les FrancoFolies de Montréal, Nuits d'Afrique, Pop Montreal, Divers/Cité, Fierté Montréal and the Montreal Fireworks Festival, Igloofest, Piknic Electronik, Montréal en Lumiere, Osheaga, Heavy Montreal, Mode + Design, Montréal complètement cirque, MUTEK, Black and Blue, and many smaller festivals. The city of Montreal is also widely recognized for its diverse and vibrant night life, which is considered a vital part of the local cultural ecosystem.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/List_of_festivals_and_parades_in_Montreal",
                                "text": "more than 100 annual festivals"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Just_for_Laughs",
                                "text": "Just for Laughs"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Les_FrancoFolies_de_Montréal",
                                "text": "Les FrancoFolies de Montréal"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Nuits_d'Afrique?action=edit&redlink=1",
                                "text": "Nuits d'Afrique"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Pop_Montreal",
                                "text": "Pop Montreal"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Divers/Cité",
                                "text": "Divers/Cité"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Fierté_Montréal",
                                "text": "Fierté Montréal"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Montreal_Fireworks_Festival",
                                "text": "Montreal Fireworks Festival"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Igloofest",
                                "text": "Igloofest"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Piknic_Électronik",
                                "text": "Piknic Electronik"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Osheaga_Festival",
                                "text": "Osheaga"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Heavy_Montréal",
                                "text": "Heavy Montreal"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/MUTEK",
                                "text": "MUTEK"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "A cultural heart of classical art and the venue for many summer festivals, the Place des Arts is a complex of different concert and theatre halls surrounding a large square in the eastern portion of downtown. Place des Arts has the headquarters of one of the world's foremost orchestras, the Montreal Symphony Orchestra . The Orchestre Métropolitain du Grand Montréal and the chamber orchestra I Musici de Montréal are two other well-regarded Montreal orchestras. Also performing at Place des Arts are the Opéra de Montréal and the city's chief ballet company Les Grands Ballets Canadiens . Internationally recognized avant-garde dance troupes such as Compagnie Marie Chouinard, La La La Human Steps, O Vertigo, and the Fondation Jean-Pierre Perreault have toured the world and worked with international popular artists on videos and concerts. The unique choreography of these troupes has paved the way for the success of the world-renowned Cirque du Soleil.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Place_des_Arts",
                                "text": "Place des Arts"
                            },
                            {
                                "url": "https://placedesarts.com/en",
                                "text": "Place des Arts"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Montreal_Symphony_Orchestra",
                                "text": "Montreal Symphony Orchestra"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Orchestre_Métropolitain_du_Grand_Montréal",
                                "text": "Orchestre Métropolitain du Grand Montréal"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/I_Musici_de_Montréal",
                                "text": "I Musici de Montréal"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Opéra_de_Montréal",
                                "text": "Opéra de Montréal"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Les_Grands_Ballets_Canadiens",
                                "text": "Les Grands Ballets Canadiens"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Compagnie_Marie_Chouinard?action=edit&redlink=1",
                                "text": "Compagnie Marie Chouinard"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/La_La_La_Human_Steps",
                                "text": "La La La Human Steps"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/O_Vertigo?action=edit&redlink=1",
                                "text": "O Vertigo"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Fondation_Jean-Pierre_Perreault?action=edit&redlink=1",
                                "text": "Fondation Jean-Pierre Perreault"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "Nicknamed la ville aux cent clochers (the city of a hundred steeples), Montreal is renowned for its churches. There are an estimated 650 churches on the island, with 450 of them dating back to the 1800s or earlier. Mark Twain noted, \"This is the first time I was ever in a city where you couldn't throw a brick without breaking a church window.\" The city has four Roman Catholic basilicas: Mary, Queen of the World Cathedral, Notre-Dame Basilica, St Patrick's Basilica, and Saint Joseph's Oratory . The Oratory is the largest church in Canada, with the second largest copper dome in the world, after Saint Peter's Basilica in Rome.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Mark_Twain",
                                "text": "Mark Twain"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Basilica",
                                "text": "basilicas"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Mary,_Queen_of_the_World_Cathedral",
                                "text": "Mary, Queen of the World Cathedral"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Notre-Dame_Basilica_(Montreal)",
                                "text": "Notre-Dame Basilica"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/St._Patrick's_Basilica,_Montreal",
                                "text": "St Patrick's Basilica"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Saint_Joseph's_Oratory",
                                "text": "Saint Joseph's Oratory"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Saint_Peter's_Basilica",
                                "text": "Saint Peter's Basilica"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "Beginning in the 1940s, Quebec literature began to shift from pastoral tales romanticising the French-Canadian countryside to writing set in the multicultural city of Montreal. Notable pioneering works describing the character of the city include Gabrielle Roy 's 1945 novel Bonheur D'Occasion, translated as The Tin Flute, and Gwethalyn Graham 's 1944 novel Earth and High Heaven . Subsequent writers of fiction who have set their work in Montreal have included Mordecai Richler, Claude Jasmin, Francine Noel, and Heather O'Neill, among many others .",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/French_Canadians",
                                "text": "French-Canadian"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Gabrielle_Roy",
                                "text": "Gabrielle Roy"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/The_Tin_Flute",
                                "text": "Bonheur D'Occasion"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Gwethalyn_Graham",
                                "text": "Gwethalyn Graham"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Mordecai_Richler",
                                "text": "Mordecai Richler"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Claude_Jasmin",
                                "text": "Claude Jasmin"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Francine_Noel",
                                "text": "Francine Noel"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Heather_O'Neill",
                                "text": "Heather O'Neill"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/List_of_people_from_Montreal",
                                "text": "many others"
                            }
                        ]
                    }
                ]
            },
            {
                "name": "Sports",
                "type": "section",
                "has_parts": [
                    {
                        "type": "paragraph",
                        "value": "The most popular sport is ice hockey . The professional hockey team, the Montreal Canadiens, is one of the Original Six teams of the National Hockey League (NHL), and has won an NHL-record 24 Stanley Cup championships. The Canadiens' most recent Stanley Cup victory came in 1993 . They have major rivalries with the Toronto Maple Leafs and Boston Bruins, both of which are also Original Six teams, and with the Ottawa Senators, the closest team geographically. The Canadiens have played at the Bell Centre since 1996. Prior to that they played at the Montreal Forum .",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Ice_hockey",
                                "text": "ice hockey"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Montreal_Canadiens",
                                "text": "Montreal Canadiens"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Original_Six",
                                "text": "Original Six"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/National_Hockey_League",
                                "text": "National Hockey League"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Stanley_Cup",
                                "text": "Stanley Cup"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/1993_Stanley_Cup_Finals",
                                "text": "1993"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Toronto_Maple_Leafs",
                                "text": "Toronto Maple Leafs"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Boston_Bruins",
                                "text": "Boston Bruins"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Ottawa_Senators",
                                "text": "Ottawa Senators"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Bell_Centre",
                                "text": "Bell Centre"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Montreal_Forum",
                                "text": "Montreal Forum"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "The Montreal Alouettes of the Canadian Football League (CFL) play at Percival Molson Memorial Stadium on the campus of McGill University for their regular-season games. Late season and playoff games are sometimes played at the much larger, enclosed Olympic Stadium, which also hosted the 2008 Grey Cup . The Alouettes have won the Grey Cup eight times, most recently in 2023 . The Alouettes have had two periods on hiatus. During the second one, the Montreal Machine played in the World League of American Football in 1991 and 1992. The McGill Redbirds, Concordia Stingers, and Université de Montréal Carabins play in the U Sports football league.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Montreal_Alouettes",
                                "text": "Montreal Alouettes"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Canadian_Football_League",
                                "text": "Canadian Football League"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Percival_Molson_Memorial_Stadium",
                                "text": "Percival Molson Memorial Stadium"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/McGill_University",
                                "text": "McGill University"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/2008_Grey_Cup",
                                "text": "2008 Grey Cup"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/111th_Grey_Cup",
                                "text": "2023"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Montreal_Machine",
                                "text": "Montreal Machine"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/World_League_of_American_Football",
                                "text": "World League of American Football"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/McGill_Redbirds_and_Martlets",
                                "text": "McGill Redbirds"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Concordia_Stingers",
                                "text": "Concordia Stingers"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Université_de_Montréal_Carabins",
                                "text": "Université de Montréal Carabins"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/U_Sports_football",
                                "text": "U Sports football"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "Montreal has a storied baseball history. The city was the home of the minor-league Montreal Royals of the International League until 1960. In 1946, Jackie Robinson broke the Baseball colour line with the Royals in an emotionally difficult year; Robinson was forever grateful for the local fans' fervent support. Major League Baseball came to town in the form of the Montreal Expos in 1969. They played their games at Jarry Park Stadium until moving into Olympic Stadium in 1977. After 36 years in Montreal, the team relocated to Washington, D.C., in 2005 and re-branded themselves as the Washington Nationals .",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Montreal_Royals",
                                "text": "Montreal Royals"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/International_League",
                                "text": "International League"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Jackie_Robinson",
                                "text": "Jackie Robinson"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Baseball_colour_line",
                                "text": "Baseball colour line"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Major_League_Baseball",
                                "text": "Major League Baseball"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Montreal_Expos",
                                "text": "Montreal Expos"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Jarry_Park_Stadium",
                                "text": "Jarry Park Stadium"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Washington_Nationals",
                                "text": "Washington Nationals"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "CF Montréal (formerly known as the Montreal Impact) are the city's professional soccer team. They play at a soccer-specific stadium called Saputo Stadium . They joined North America's biggest soccer league, Major League Soccer, in 2012. The Montreal games of the 2007 FIFA U-20 World Cup and 2014 FIFA U-20 Women's World Cup were held at Olympic Stadium, and the venue hosted Montreal games in the 2015 FIFA Women's World Cup .",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/CF_Montréal",
                                "text": "CF Montréal"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Soccer-specific_stadium",
                                "text": "soccer-specific stadium"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Saputo_Stadium",
                                "text": "Saputo Stadium"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Major_League_Soccer",
                                "text": "Major League Soccer"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/2007_FIFA_U-20_World_Cup",
                                "text": "2007 FIFA U-20 World Cup"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/2014_FIFA_U-20_Women's_World_Cup",
                                "text": "2014 FIFA U-20 Women's World Cup"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/2015_FIFA_Women's_World_Cup",
                                "text": "2015 FIFA Women's World Cup"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "Montreal is the site of a high-profile auto racing event each year: the Canadian Grand Prix of Formula One (F1) racing. This race takes place on the Circuit Gilles Villeneuve on Île Notre-Dame. In 2009, the race was dropped from the Formula One calendar, to the chagrin of some fans, but the Canadian Grand Prix returned to the Formula One calendar in 2010. It was dropped from the calendar again in 2020 and 2021, due to COVID-19 pandemic, but racing resumed in 2022, with the 2022 Canadian Grand Prix . The Circuit Gilles Villeneuve also hosted a round of the Champ Car World Series from 2002 to 2007, and was home to the NAPA Auto Parts 200, a NASCAR Nationwide Series race, and the Montréal 200, a Grand Am Rolex Sports Car Series race.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Auto_racing",
                                "text": "auto racing"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Circuit_Gilles_Villeneuve",
                                "text": "Circuit Gilles Villeneuve"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/COVID-19_pandemic",
                                "text": "COVID-19 pandemic"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/2022_Canadian_Grand_Prix",
                                "text": "2022 Canadian Grand Prix"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Champ_Car_World_Series",
                                "text": "Champ Car World Series"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/NAPA_Auto_Parts_200",
                                "text": "NAPA Auto Parts 200"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/NASCAR_Nationwide_Series",
                                "text": "NASCAR Nationwide Series"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Grand_American_Road_Racing_Association",
                                "text": "Grand Am"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Rolex_Sports_Car_Series",
                                "text": "Rolex Sports Car Series"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "Uniprix Stadium, built in 1993 on the site of Jarry Park, is used for the National Bank Open (formerly known as the Rogers Cup) men's and women's tennis tournaments. The men's tournament is a Masters 1000 event on the ATP Tour, and the women's tournament is a Premier tournament on the WTA Tour . The men's and women's tournaments alternate between Montreal and Toronto every year.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Uniprix_Stadium",
                                "text": "Uniprix Stadium"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Canadian_Open_(tennis)",
                                "text": "National Bank Open"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Masters_1000",
                                "text": "Masters 1000"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/ATP_Tour",
                                "text": "ATP Tour"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/WTA_Premier_tournaments",
                                "text": "Premier tournament"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/WTA_Tour",
                                "text": "WTA Tour"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "Montreal was the host of the 1976 Summer Olympic Games. The stadium cost $1.5 billion; with interest that figure ballooned to nearly $3 billion, and was paid off in December 2006. Montreal also hosted the first ever World Outgames in the summer of 2006, attracting over 16,000 participants engaged in 35 sporting activities.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/World_Outgames",
                                "text": "World Outgames"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "Montreal was the host city for the 17th unicycling world championship and convention (UNICON) in August 2014.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/UNICON",
                                "text": "UNICON"
                            }
                        ]
                    }
                ]
            },
            {
                "name": "Media",
                "type": "section",
                "has_parts": [
                    {
                        "type": "paragraph",
                        "value": "Montreal is Canada's second-largest media market, and the centre of Canada's francophone media industry."
                    },
                    {
                        "type": "paragraph",
                        "value": "There are four over-the-air English-language television stations: CBMT-DT (CBC Television), CFCF-DT (CTV), CKMI-DT (Global) and CJNT-DT (Citytv). There are also five over-the-air French-language television stations: CBFT-DT (Ici Radio-Canada), CFTM-DT (TVA), CFJP-DT (Noovo), CIVM-DT (Télé-Québec), and CFTU-DT (Canal Savoir).",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Terrestrial_television",
                                "text": "over-the-air"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/CBMT-DT",
                                "text": "CBMT-DT"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/CBC_Television",
                                "text": "CBC Television"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/CFCF-DT",
                                "text": "CFCF-DT"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/CTV_Television_Network",
                                "text": "CTV"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/CKMI-DT",
                                "text": "CKMI-DT"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Global_Television_Network",
                                "text": "Global"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/CJNT-DT",
                                "text": "CJNT-DT"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Citytv",
                                "text": "Citytv"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/CBFT-DT",
                                "text": "CBFT-DT"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Ici_Radio-Canada_Télé",
                                "text": "Ici Radio-Canada"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/CFTM-DT",
                                "text": "CFTM-DT"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/TVA_(Canadian_TV_network)",
                                "text": "TVA"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/CFJP-DT",
                                "text": "CFJP-DT"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Noovo",
                                "text": "Noovo"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Télé-Québec",
                                "text": "Télé-Québec"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/CFTU-DT",
                                "text": "CFTU-DT"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "Montreal has three daily newspapers, the English-language Montreal Gazette and the French-language Le Journal de Montréal, and Le Devoir; another French-language daily, La Presse, became an online daily in 2018. There are two free French dailies, Métro and 24 Heures . Montreal has numerous weekly tabloids and community newspapers serving various neighbourhoods, ethnic groups and schools.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Montreal_Gazette",
                                "text": "Montreal Gazette"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Le_Journal_de_Montréal",
                                "text": "Le Journal de Montréal"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Le_Devoir",
                                "text": "Le Devoir"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/La_Presse_(Canadian_newspaper)",
                                "text": "La Presse"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Metro_International",
                                "text": "Métro"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/24_Hours_(newspaper)",
                                "text": "24 Heures"
                            }
                        ]
                    }
                ]
            },
            {
                "name": "Government",
                "type": "section",
                "has_parts": [
                    {
                        "type": "paragraph",
                        "value": "The head of the city government in Montreal is the mayor, who is first among equals in the city council.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/City_government_in_Montreal",
                                "text": "city government in Montreal"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/First_among_equals",
                                "text": "first among equals"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "The city council is a democratically elected institution and is the final decision-making authority in the city, although much power is centralized in the executive committee. The council consists of 65 members from all boroughs. The council has jurisdiction over many matters, including public security, agreements with other governments, subsidy programs, the environment, urban planning, and a three-year capital expenditure program. The council is required to supervise, standardize or approve certain decisions made by the borough councils.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Natural_environment",
                                "text": "environment"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Urban_planning",
                                "text": "urban planning"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "Reporting directly to the council, the executive committee exercises decision-making powers similar to those of the cabinet in a parliamentary system and is responsible for preparing various documents including budgets and by-laws, submitted to the council for approval. The decision-making powers of the executive committee cover, in particular, the awarding of contracts or grants, the management of human and financial resources, supplies and buildings. It may also be assigned further powers by the city council.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Parliamentary_system",
                                "text": "parliamentary system"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Budget",
                                "text": "budgets"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/By-law",
                                "text": "by-laws"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "Standing committees are the prime instruments for public consultation. They are responsible for the public study of pending matters and for making the appropriate recommendations to the council. They also review the annual budget forecasts for departments under their jurisdiction. A public notice of meeting is published in both French and English daily newspapers at least seven days before each meeting. All meetings include a public question period. The standing committees, of which there are seven, have terms lasting two years. In addition, the City Council may decide to create special committees at any time. Each standing committee is made up of seven to nine members, including a chairman and a vice-chairman. The members are all elected municipal officers, with the exception of a representative of the government of Quebec on the public security committee."
                    },
                    {
                        "type": "paragraph",
                        "value": "The city is only one component of the larger Montreal Metropolitan Community (Communauté Métropolitaine de Montréal, CMM), which is in charge of planning, coordinating, and financing economic development, public transportation, garbage collection and waste management, etc., across the metropolitan area. The president of the CMM is the mayor of Montreal. The CMM covers 4,360 km (1,680 sq mi), with 3.6 million inhabitants in 2006.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Waste_management",
                                "text": "waste management"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "Montreal is the seat of the judicial district of Montreal, which includes the city and the other communities on the island.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Judicial_districts_of_Quebec",
                                "text": "judicial district"
                            }
                        ]
                    }
                ]
            },
            {
                "name": "Policing",
                "type": "section",
                "has_parts": [
                    {
                        "type": "paragraph",
                        "value": "Law enforcement on the island itself is provided by the Service de Police de la Ville de Montréal, or the SPVM for short.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/SPVM",
                                "text": "Service de Police de la Ville de Montréal"
                            }
                        ]
                    }
                ]
            },
            {
                "name": "Crime",
                "type": "section",
                "has_parts": [
                    {
                        "type": "paragraph",
                        "value": "Since 1975, when Montreal's homicide rate peaked at around 10.3 per 100,000 people with a total of 112 murders, the overall crime rate in Montreal has declined, with a few notable exceptions, reaching a minimum in 2016 with 23 murders. Sex crimes have increased 14.5 per cent between 2015 and 2016 and fraud cases have increased by 13 per cent over the same period. The major criminal organizations active in Montreal are the Rizzuto crime family, Hells Angels and West End Gang . However, in the 2020s, the city has seen an increase in overall crime, with a notable increase in homicides. 25 homicides were reported in 2020 which matched the number reported in 2019. The next year saw a 48% increase in murders with a total of 37 in 2021, giving the city a homicide rate of around 2.1 per 100,000 people. The Montreal Police Annual Report for 2021 showed that there were 144 shootings across the city, or an average of one shooting every 2.5 days. In comparison, there were 71 shootings recorded the year before. 2022 saw another 10.8% increase in homicides, with a total of 41 being reported (giving a slightly higher homicide rate of 2.3 per 100,000 people), the highest number since 2007, when there were 42.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Rizzuto_crime_family",
                                "text": "Rizzuto crime family"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Hells_Angels",
                                "text": "Hells Angels"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/West_End_Gang",
                                "text": "West End Gang"
                            }
                        ]
                    }
                ]
            },
            {
                "name": "Education",
                "type": "section",
                "has_parts": [
                    {
                        "type": "paragraph",
                        "value": "The education system in Quebec is different from other systems in North America. Between high school (which ends at grade 11) and university, students must go through an additional school called CEGEP . CEGEPs offer pre-university (2-years) and technical (3-years) programs. In Montreal, seventeen CEGEPs offer courses in French and five in English.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Education_in_Quebec",
                                "text": "education system in Quebec"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/CEGEP",
                                "text": "CEGEP"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/List_of_CEGEPs",
                                "text": "seventeen CEGEPs"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "French-language elementary and secondary public schools in Montreal are operated by the Centre de services scolaire de Montréal (CSSDM), Centre de services scolaire Marguerite-Bourgeoys and the Centre de services scolaire de la Pointe-de-l'Île .",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Centre_de_services_scolaire_de_Montréal",
                                "text": "Centre de services scolaire de Montréal"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Centre_de_services_scolaire_Marguerite-Bourgeoys",
                                "text": "Centre de services scolaire Marguerite-Bourgeoys"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Centre_de_services_scolaire_de_la_Pointe-de-l'Île",
                                "text": "Centre de services scolaire de la Pointe-de-l'Île"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "English-language elementary and secondary public schools on Montreal Island are operated by the English Montreal School Board and the Lester B. Pearson School Board .",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/English_Montreal_School_Board",
                                "text": "English Montreal School Board"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Lester_B._Pearson_School_Board",
                                "text": "Lester B. Pearson School Board"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "With four universities, ten other degree-awarding institutions, and 12 CEGEPs in an 8 km (5.0 mi) radius, Montreal has the highest concentration of post-secondary students of all major cities in North America (4.38 students per 100 residents, followed by Boston at 4.37 students per 100 residents).",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/CEGEP",
                                "text": "CEGEPs"
                            }
                        ]
                    },
                    {
                        "name": "Higher education (English)",
                        "type": "section",
                        "has_parts": [
                            {
                                "type": "list",
                                "has_parts": [
                                    {
                                        "type": "list_item",
                                        "value": "McGill University is one of Canada's leading post-secondary institutions and is widely regarded as a world-class institution. In 2021, McGill was ranked as the top medical-doctoral university in Canada for the seventeenth consecutive year by Maclean's and second in Canada and the 27th best university in the world by the QS World University Rankings .",
                                        "links": [
                                            {
                                                "url": "https://en.wikipedia.org/wiki/McGill_University",
                                                "text": "McGill University"
                                            },
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Maclean's",
                                                "text": "Maclean's"
                                            },
                                            {
                                                "url": "https://en.wikipedia.org/wiki/QS_World_University_Rankings",
                                                "text": "QS World University Rankings"
                                            }
                                        ]
                                    },
                                    {
                                        "type": "list_item",
                                        "value": "Concordia University was created from the merger of Sir George Williams University and Loyola College in 1974. The university has been ranked as one of the top comprehensive universities in Canada by Macleans.",
                                        "links": [
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Concordia_University_(Montreal)",
                                                "text": "Concordia University"
                                            },
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Concordia_University_(Montreal)#Sir_George_Williams_University",
                                                "text": "Sir George Williams University"
                                            },
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Concordia_University_(Montreal)#Loyola_College",
                                                "text": "Loyola College"
                                            }
                                        ]
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "name": "Higher education (French)",
                        "type": "section",
                        "has_parts": [
                            {
                                "type": "list",
                                "has_parts": [
                                    {
                                        "type": "list_item",
                                        "value": "Université de Montréal (UdeM) is the second largest research university in Canada and ranked as one of the top universities in Canada. Two separate institutions are affiliated to the university: the École Polytechnique de Montréal (School of Engineering) and HEC Montréal (School of Business). HEC Montreal was founded in 1907 and is considered one of the best business schools in Canada.",
                                        "links": [
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Université_de_Montréal",
                                                "text": "Université de Montréal"
                                            },
                                            {
                                                "url": "https://en.wikipedia.org/wiki/École_Polytechnique_de_Montréal",
                                                "text": "École Polytechnique de Montréal"
                                            },
                                            {
                                                "url": "https://en.wikipedia.org/wiki/HEC_Montréal",
                                                "text": "HEC Montréal"
                                            }
                                        ]
                                    },
                                    {
                                        "type": "list_item",
                                        "value": "Université du Québec à Montréal (UQAM) is the Montreal campus of Université du Québec . UQAM generally specializes in liberal-arts, although many programs related to the sciences are available.The Université du Québec network also has three separately run schools in Montreal, notably the École de technologie supérieure (ETS), the École nationale d'administration publique (ÉNAP) and the Institut national de la recherche scientifique (INRS) .",
                                        "links": [
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Université_du_Québec_à_Montréal",
                                                "text": "Université du Québec à Montréal"
                                            },
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Université_du_Québec",
                                                "text": "Université du Québec"
                                            },
                                            {
                                                "url": "https://en.wikipedia.org/wiki/École_de_technologie_supérieure",
                                                "text": "École de technologie supérieure"
                                            },
                                            {
                                                "url": "https://en.wikipedia.org/wiki/École_nationale_d'administration_publique",
                                                "text": "École nationale d'administration publique"
                                            },
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Institut_national_de_la_recherche_scientifique",
                                                "text": "Institut national de la recherche scientifique"
                                            }
                                        ]
                                    },
                                    {
                                        "type": "list_item",
                                        "value": "The Université du Québec network also has three separately run schools in Montreal, notably the École de technologie supérieure (ETS), the École nationale d'administration publique (ÉNAP) and the Institut national de la recherche scientifique (INRS) .",
                                        "links": [
                                            {
                                                "url": "https://en.wikipedia.org/wiki/École_de_technologie_supérieure",
                                                "text": "École de technologie supérieure"
                                            },
                                            {
                                                "url": "https://en.wikipedia.org/wiki/École_nationale_d'administration_publique",
                                                "text": "École nationale d'administration publique"
                                            },
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Institut_national_de_la_recherche_scientifique",
                                                "text": "Institut national de la recherche scientifique"
                                            }
                                        ]
                                    },
                                    {
                                        "type": "list_item",
                                        "value": "L'Institut de formation théologique de Montréal des Prêtres de Saint-Sulpice (IFTM) specializes in theology and philosophy.",
                                        "links": [
                                            {
                                                "url": "https://en.wikipedia.org/wiki/L'Institut_de_formation_théologique_de_Montréal?action=edit&redlink=1",
                                                "text": "L'Institut de formation théologique de Montréal"
                                            }
                                        ]
                                    },
                                    {
                                        "type": "list_item",
                                        "value": "Institut d'hôtellerie et de tourisme du Québec (IHTQ) offers an Applied Bachelor in Hospitality and Hotel Management.",
                                        "links": [
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Institut_d'hôtellerie_et_de_tourisme_du_Québec?action=edit&redlink=1",
                                                "text": "Institut d'hôtellerie et de tourisme du Québec"
                                            }
                                        ]
                                    },
                                    {
                                        "type": "list_item",
                                        "value": "Conservatoire de musique du Québec à Montréal offers both a Bachelor and a Master program in classical music.",
                                        "links": [
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Conservatoire_de_musique_du_Québec_à_Montréal",
                                                "text": "Conservatoire de musique du Québec à Montréal"
                                            },
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Bachelor's_degree",
                                                "text": "Bachelor"
                                            },
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Master's_degree",
                                                "text": "Master"
                                            }
                                        ]
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "Additionally, two French-language universities, Université de Sherbrooke and Université Laval have campuses in the nearby suburb of Longueuil on Montreal's south shore . Also, l'Institut de pastorale des Dominicains is Montreal's university centre of Ottawa's Collège Universitaire Dominicain /Dominican University College . The Faculté de théologie évangélique is Nova Scotia 's Acadia University Montreal based serving French Protestant community in Canada by offering both a Bachelor and a Master program in theology",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Université_de_Sherbrooke",
                                        "text": "Université de Sherbrooke"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Université_Laval",
                                        "text": "Université Laval"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/South_Shore_(Montreal)",
                                        "text": "south shore"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Institut_de_pastorale_des_Dominicains",
                                        "text": "Institut de pastorale des Dominicains"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Collège_Universitaire_Dominicain",
                                        "text": "Collège Universitaire Dominicain"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Dominican_University_College",
                                        "text": "Dominican University College"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Faculté_de_théologie_évangélique?action=edit&redlink=1",
                                        "text": "Faculté de théologie évangélique"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Nova_Scotia",
                                        "text": "Nova Scotia"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Acadia_University",
                                        "text": "Acadia University"
                                    }
                                ]
                            }
                        ]
                    }
                ]
            },
            {
                "name": "Transportation",
                "type": "section",
                "has_parts": [
                    {
                        "type": "paragraph",
                        "value": "Like many major cities, Montreal has a problem with vehicular traffic congestion. Commuting traffic from the cities and towns in the West Island (such as Dollard-des-Ormeaux and Pointe-Claire) is compounded by commuters entering the city that use twenty-four road crossings from numerous off-island suburbs on the North and South Shores. The width of the Saint Lawrence River has made the construction of fixed links to the south shore expensive and difficult. There are presently four road bridges (including two of the country's busiest) along with one bridge-tunnel, two railway bridges, and a metro line. The far narrower Rivière des Prairies to the city's north, separating Montreal from Laval, is spanned by nine road bridges (seven to the city of Laval and two that span directly to the north shore) and a Metro line.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Dollard-des-Ormeaux",
                                "text": "Dollard-des-Ormeaux"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Pointe-Claire",
                                "text": "Pointe-Claire"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/North_Shore_(Laval)",
                                "text": "North"
                            }
                        ]
                    },
                    {
                        "type": "paragraph",
                        "value": "The island of Montreal is a hub for the Quebec Autoroute system, and is served by Quebec Autoroutes A-10 (known as the Bonaventure Expressway on the island of Montreal), A-15 (aka the Décarie Expressway south of the A-40 and the Laurentian Autoroute to the north of it), A-13 (aka Chomedey Autoroute), A-20, A-25, A-40 (part of the Trans-Canada Highway system, and known as \"The Metropolitan\" or simply \"The Met\" in its elevated mid-town section), A-520 and R-136 (aka the Ville-Marie Autoroute). Many of these Autoroutes are frequently congested at rush hour . However, in recent years, the government has acknowledged this problem and is working on long-term solutions to alleviate the congestion. One such example is the extension of Quebec Autoroute 30 on Montreal's south shore, which will be a bypass for trucks and intercity traffic.",
                        "links": [
                            {
                                "url": "https://en.wikipedia.org/wiki/Autoroute_(Quebec)",
                                "text": "Autoroute"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Quebec_Autoroute_10",
                                "text": "A-10"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Quebec_Autoroute_15",
                                "text": "A-15"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Quebec_Autoroute_13",
                                "text": "A-13"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Quebec_Autoroute_20",
                                "text": "A-20"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Quebec_Autoroute_25",
                                "text": "A-25"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Quebec_Autoroute_40",
                                "text": "A-40"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Trans-Canada_Highway",
                                "text": "Trans-Canada Highway"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Quebec_Autoroute_520",
                                "text": "A-520"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Quebec_Route_136_(Montreal)",
                                "text": "R-136"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Rush_hour",
                                "text": "rush hour"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Quebec_Autoroute_30",
                                "text": "Quebec Autoroute 30"
                            },
                            {
                                "url": "https://en.wikipedia.org/wiki/Bypass_(road)",
                                "text": "bypass"
                            }
                        ]
                    },
                    {
                        "name": "Société de transport de Montréal",
                        "type": "section",
                        "has_parts": [
                            {
                                "type": "paragraph",
                                "value": "Public local transport is served by a network of buses, subways, and commuter trains that extend across and off the island. The subway and bus system are operated by STM (Société de transport de Montréal, “Montreal Transit Company”). The STM bus network consists of 203 daytime and 23 night time routes. STM bus routes serve 1,347,900 passengers on an average weekday in 2010. It also provides adapted transport and wheelchair-accessible buses. The STM won the award of Outstanding Public Transit System in North America by the APTA in 2010. It was the first time a Canadian company won this prize.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/List_of_Montreal_bus_routes",
                                        "text": "STM bus network"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "The Metro was inaugurated in 1966 and has 68 stations on four lines. Total daily passengers is 1,050,800 passengers on an average weekday (as of Q1 2010). Each station was designed by different architects with individual themes and features original artwork, and the trains run on rubber tires, making the system quieter than most. The project was initiated by Montreal Mayor Jean Drapeau, who later brought the Summer Olympic Games to Montreal in 1976. The Metro system has long had a station on the South Shore in Longueuil, and in 2007 was extended to the city of Laval, north of Montreal, with three new stations. The metro has recently been modernizing its trains, purchasing new Azur models with inter-connected wagons.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Jean_Drapeau",
                                        "text": "Jean Drapeau"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Longueuil,_Quebec",
                                        "text": "Longueuil"
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "name": "Air",
                        "type": "section",
                        "has_parts": [
                            {
                                "type": "paragraph",
                                "value": "Montreal has two international airports, one for passengers only, the other for cargo. Pierre Elliott Trudeau International Airport (also known as Dorval Airport) in the City of Dorval serves all commercial passenger traffic and is the headquarters of Air Canada and Air Transat. To the north of the city is Montreal Mirabel International Airport in Mirabel, which was envisioned as Montreal's primary airport but which now serves cargo flights along with MEDEVACs and general aviation and some passenger services. In 2018, Trudeau was the third busiest airport in Canada by passenger traffic and aircraft movements, handling 19.42 million passengers, and 240,159 aircraft movements. With 63% of its passengers being on non-domestic flights it has the largest percentage of international flights of any Canadian airport.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Montréal-Pierre_Elliott_Trudeau_International_Airport",
                                        "text": "Pierre Elliott Trudeau International Airport"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Dorval",
                                        "text": "Dorval"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Montréal-Mirabel_International_Airport",
                                        "text": "Montreal Mirabel International Airport"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Mirabel,_Quebec",
                                        "text": "Mirabel"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/MEDEVAC",
                                        "text": "MEDEVACs"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/General_aviation",
                                        "text": "general aviation"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/List_of_the_busiest_airports_in_Canada",
                                        "text": "third busiest airport in Canada"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "It is one of Air Canada's major hubs and operates on average approximately 2,400 flights per week between Montreal and 155 destinations, spread on five continents .",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Airline_hub",
                                        "text": "hubs"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Continents",
                                        "text": "continents"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "Airlines servicing Trudeau offer year-round non-stop flights to five continents, namely Africa, Asia, Europe, North America and South America. It is one of only two airports in Canada with direct flights to five continents or more.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Non-stop_flight",
                                        "text": "non-stop flights"
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "name": "Rail",
                        "type": "section",
                        "has_parts": [
                            {
                                "type": "paragraph",
                                "value": "Montreal-based Via Rail Canada provides rail service to other cities in Canada, particularly to Quebec City and Toronto along the Quebec City – Windsor Corridor . Amtrak, the U.S. national passenger rail system, operates its Adirondack daily to New York. All intercity trains and most commuter trains operate out of Central Station .",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Quebec_City_–_Windsor_Corridor",
                                        "text": "Quebec City – Windsor Corridor"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Amtrak",
                                        "text": "Amtrak"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Adirondack_(Amtrak)",
                                        "text": "Adirondack"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Commuter_train",
                                        "text": "commuter trains"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Central_Station_(Montreal)",
                                        "text": "Central Station"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "Canadian Pacific Railway (CPR) was founded here in 1881. Its corporate headquarters occupied Windsor Station at 910 Peel Street until 1995, when it moved to Calgary, Alberta. With the Port of Montreal kept open year-round by icebreakers, lines to Eastern Canada became surplus, and now Montreal is the eastern and intermodal freight terminus of CPR's successor company, Canadian Pacific Kansas City (CPKC). CPKC connects at Montreal with the Port of Montreal, the Delaware and Hudson Railway to New York, the Quebec Gatineau Railway to Quebec City and Buckingham, the Central Maine and Quebec Railway to Halifax, and Canadian National Railway (CN). The CPR's flagship train, The Canadian, ran daily from Windsor Station to Vancouver, but in 1978 all passenger services were transferred to Via. Since 1990, The Canadian has terminated in Toronto instead of in Montreal.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Windsor_Station_(Montreal)",
                                        "text": "Windsor Station"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Peel_Street,_Montreal",
                                        "text": "Peel Street"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Calgary",
                                        "text": "Calgary"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Canadian_Pacific_Kansas_City",
                                        "text": "Canadian Pacific Kansas City"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Delaware_and_Hudson_Railway",
                                        "text": "Delaware and Hudson Railway"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Quebec_Gatineau_Railway",
                                        "text": "Quebec Gatineau Railway"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Buckingham,_Quebec",
                                        "text": "Buckingham"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Central_Maine_and_Quebec_Railway",
                                        "text": "Central Maine and Quebec Railway"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Canadian_National_Railway",
                                        "text": "Canadian National Railway"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/The_Canadian",
                                        "text": "The Canadian"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Vancouver",
                                        "text": "Vancouver"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "Montreal-based CN was formed in 1919 by the Canadian government following a series of country-wide rail bankruptcies. It was formed from the Grand Trunk, Midland and Canadian Northern Railways, and has risen to become CPR's chief rival in freight carriage in Canada. Like the CPR, CN divested itself of passenger services in favour of Via. CN's flagship train, the Super Continental, ran daily from Central Station to Vancouver and subsequently became a Via train in 1978. It was eliminated in 1990 in favour of rerouting The Canadian .",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Grand_Trunk_Railway",
                                        "text": "Grand Trunk"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Canadian_Northern_Railway",
                                        "text": "Canadian Northern Railways"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Super_Continental",
                                        "text": "Super Continental"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "The commuter rail system is managed and operated by Exo, and reaches the outlying areas of Greater Montreal with six lines. It carried an average of 79,000 daily passengers in 2014, making it the seventh busiest in North America following New York, Chicago, Toronto, Boston, Philadelphia, and Mexico City.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Exo_(public_transit)",
                                        "text": "Exo"
                                    }
                                ]
                            },
                            {
                                "type": "paragraph",
                                "value": "On April 22, 2016, the forthcoming automated rapid transit system, the Réseau express métropolitain (REM), was unveiled. Groundbreaking occurred April 12, 2018, and construction of the 67-kilometre-long (42 mi) network – consisting of three branches, 26 stations, and the conversion of the region's busiest commuter railway – commenced the following month. To be opened in three phases as of 2022, the REM will be completed by mid-2024, becoming the fourth largest automated rapid transit network after the Dubai Metro, the Singapore Mass Rapid Transit, and the Vancouver SkyTrain . Most of it will be financed by pension fund manager Caisse de dépôt et placement du Québec (CDPQ Infra).",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Rapid_transit",
                                        "text": "rapid transit"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Réseau_express_métropolitain",
                                        "text": "Réseau express métropolitain"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Dubai_Metro",
                                        "text": "Dubai Metro"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Mass_Rapid_Transit_(Singapore)",
                                        "text": "Singapore Mass Rapid Transit"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/SkyTrain_(Vancouver)",
                                        "text": "Vancouver SkyTrain"
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "name": "Bike Share Program",
                        "type": "section",
                        "has_parts": [
                            {
                                "type": "paragraph",
                                "value": "The city of Montreal is world-renowned for being in the top 20 most cyclist-friendly cities around the globe. It follows that they have one of the world's most successful bike share systems in BIXI. First launched in 2009 with Montreal-based PBSC Urban Solutions ICONIC bikes, the bicycle-sharing scheme has since grown its fleet to include 750 docking and charging stations across the different neighbourhoods with 9000 bikes available for users. In what the STM states is a mission to combine different forms of mobility, transit card holders can now take advantage of their membership to also rent bicycles at select stations.",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/BIXI_Montréal",
                                        "text": "BIXI."
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Société_de_transport_de_Montréal",
                                        "text": "STM"
                                    }
                                ]
                            }
                        ]
                    }
                ]
            },
            {
                "name": "Notable people",
                "type": "section"
            },
            {
                "name": "International relations",
                "type": "section",
                "has_parts": [
                    {
                        "name": "Sister cities",
                        "type": "section",
                        "has_parts": [
                            {
                                "type": "list",
                                "has_parts": [
                                    {
                                        "type": "list_item",
                                        "value": "Algiers, Algeria – 1999",
                                        "links": [
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Algiers",
                                                "text": "Algiers"
                                            }
                                        ]
                                    },
                                    {
                                        "type": "list_item",
                                        "value": "Brussels, Belgium",
                                        "links": [
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Brussels",
                                                "text": "Brussels"
                                            }
                                        ]
                                    },
                                    {
                                        "type": "list_item",
                                        "value": "Bucharest, Romania",
                                        "links": [
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Bucharest",
                                                "text": "Bucharest"
                                            }
                                        ]
                                    },
                                    {
                                        "type": "list_item",
                                        "value": "Busan, South Korea – 2000",
                                        "links": [
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Busan",
                                                "text": "Busan"
                                            }
                                        ]
                                    },
                                    {
                                        "type": "list_item",
                                        "value": "Boston, United States – 1995",
                                        "links": [
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Boston",
                                                "text": "Boston"
                                            }
                                        ]
                                    },
                                    {
                                        "type": "list_item",
                                        "value": "Guadalajara, Mexico – 2004",
                                        "links": [
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Guadalajara",
                                                "text": "Guadalajara"
                                            }
                                        ]
                                    },
                                    {
                                        "type": "list_item",
                                        "value": "Hanoi, Vietnam – 1997",
                                        "links": [
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Hanoi",
                                                "text": "Hanoi"
                                            }
                                        ]
                                    },
                                    {
                                        "type": "list_item",
                                        "value": "Hiroshima, Japan – 1998",
                                        "links": [
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Hiroshima",
                                                "text": "Hiroshima"
                                            }
                                        ]
                                    },
                                    {
                                        "type": "list_item",
                                        "value": "Lyon, France – 1979",
                                        "links": [
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Lyon",
                                                "text": "Lyon"
                                            }
                                        ]
                                    },
                                    {
                                        "type": "list_item",
                                        "value": "Manila, Philippines – 2005",
                                        "links": [
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Manila",
                                                "text": "Manila"
                                            }
                                        ]
                                    },
                                    {
                                        "type": "list_item",
                                        "value": "Melbourne, Australia – 2007",
                                        "links": [
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Melbourne",
                                                "text": "Melbourne"
                                            }
                                        ]
                                    },
                                    {
                                        "type": "list_item",
                                        "value": "Port-au-Prince, Haiti – 1995",
                                        "links": [
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Port-au-Prince",
                                                "text": "Port-au-Prince"
                                            }
                                        ]
                                    },
                                    {
                                        "type": "list_item",
                                        "value": "Quito, Ecuador – 1997",
                                        "links": [
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Quito",
                                                "text": "Quito"
                                            }
                                        ]
                                    },
                                    {
                                        "type": "list_item",
                                        "value": "Rio de Janeiro, Brazil – 1998",
                                        "links": [
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Rio_de_Janeiro",
                                                "text": "Rio de Janeiro"
                                            }
                                        ]
                                    },
                                    {
                                        "type": "list_item",
                                        "value": "San Salvador, El Salvador – 2001",
                                        "links": [
                                            {
                                                "url": "https://en.wikipedia.org/wiki/San_Salvador",
                                                "text": "San Salvador"
                                            }
                                        ]
                                    },
                                    {
                                        "type": "list_item",
                                        "value": "Shanghai, China – 1985",
                                        "links": [
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Shanghai",
                                                "text": "Shanghai"
                                            }
                                        ]
                                    },
                                    {
                                        "type": "list_item",
                                        "value": "Tunis, Tunisia – 1999",
                                        "links": [
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Tunis",
                                                "text": "Tunis"
                                            }
                                        ]
                                    },
                                    {
                                        "type": "list_item",
                                        "value": "Yerevan, Armenia – 1998",
                                        "links": [
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Yerevan",
                                                "text": "Yerevan"
                                            }
                                        ]
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "name": "Friendship cities",
                        "type": "section",
                        "has_parts": [
                            {
                                "type": "list",
                                "has_parts": [
                                    {
                                        "type": "list_item",
                                        "value": "Paris, France – 2006",
                                        "links": [
                                            {
                                                "url": "https://en.wikipedia.org/wiki/Paris",
                                                "text": "Paris"
                                            }
                                        ]
                                    }
                                ]
                            }
                        ]
                    }
                ]
            },
            {
                "name": "See also",
                "type": "section",
                "has_parts": [
                    {
                        "type": "list",
                        "has_parts": [
                            {
                                "type": "list_item",
                                "value": "List of anglophone communities in Quebec",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/List_of_anglophone_communities_in_Quebec",
                                        "text": "List of anglophone communities in Quebec"
                                    }
                                ]
                            },
                            {
                                "type": "list_item",
                                "value": "List of mayors of Montreal",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/List_of_mayors_of_Montreal",
                                        "text": "List of mayors of Montreal"
                                    }
                                ]
                            },
                            {
                                "type": "list_item",
                                "value": "List of Montreal music venues",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/List_of_Montreal_music_venues",
                                        "text": "List of Montreal music venues"
                                    }
                                ]
                            },
                            {
                                "type": "list_item",
                                "value": "List of shopping malls in Montreal",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/List_of_shopping_malls_in_Montreal",
                                        "text": "List of shopping malls in Montreal"
                                    }
                                ]
                            },
                            {
                                "type": "list_item",
                                "value": "List of tallest buildings in Montreal",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/List_of_tallest_buildings_in_Montreal",
                                        "text": "List of tallest buildings in Montreal"
                                    }
                                ]
                            },
                            {
                                "type": "list_item",
                                "value": "Montreal International Games Summit",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Montreal_International_Games_Summit",
                                        "text": "Montreal International Games Summit"
                                    }
                                ]
                            },
                            {
                                "type": "list_item",
                                "value": "Order of Montreal",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Order_of_Montreal",
                                        "text": "Order of Montreal"
                                    }
                                ]
                            },
                            {
                                "type": "list_item",
                                "value": "Royal eponyms in Canada",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Royal_eponyms_in_Canada",
                                        "text": "Royal eponyms in Canada"
                                    }
                                ]
                            }
                        ]
                    }
                ]
            },
            {
                "name": "Further reading",
                "type": "section",
                "has_parts": [
                    {
                        "type": "list",
                        "has_parts": [
                            {
                                "type": "list_item",
                                "value": "Collard, Edgar A. (1976). Montréal: The Days That Are No More, in series, Totem Book. This ed. slightly edited . Toronto, Ont.: Doubleday Canada,, cop. 1976. x, 140, p., ill. in b&w with maps and numerous sketches. ISBN 0-00-216686-0 .",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/ISBN_(identifier)",
                                        "text": "ISBN"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Special:BookSources/0-00-216686-0",
                                        "text": "0-00-216686-0"
                                    }
                                ]
                            },
                            {
                                "type": "list_item",
                                "value": "Gagnon, Robert (1996). Anglophones at the C.E.C.M.: a Reflection of the Linguistic Duality of Montréal . Trans. by Peter Keating. Montréal: Commission des écoles catholiques de Montréal. 124 p., ill. with b&w photos. ISBN 2-920855-98-0 .",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/ISBN_(identifier)",
                                        "text": "ISBN"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Special:BookSources/2-920855-98-0",
                                        "text": "2-920855-98-0"
                                    }
                                ]
                            },
                            {
                                "type": "list_item",
                                "value": "Harris, David; Lyon, Patricia (2004). Montréal . Fodor's. ISBN 978-1-4000-1315-9 . Archived from the original on October 18, 2023. Retrieved December 28, 2021 .",
                                "links": [
                                    {
                                        "url": "https://books.google.com/books?id=bpxmNNnjBQoC&pg=PP1",
                                        "text": "Montréal"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/ISBN_(identifier)",
                                        "text": "ISBN"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Special:BookSources/978-1-4000-1315-9",
                                        "text": "978-1-4000-1315-9"
                                    },
                                    {
                                        "url": "https://web.archive.org/web/20231018183728/https://books.google.com/books?id=bpxmNNnjBQoC&pg=PP1#v=onepage&q&f=false",
                                        "text": "Archived"
                                    }
                                ]
                            },
                            {
                                "type": "list_item",
                                "value": "Heritage Montréal (1992). Steps in Time = Patrimoine en marche . Montréal: Québécor. 4 vol. of 20, 20 p. each. Text printed \"tête-bêche\" in English and in French. On title covers: \"Montréal, fête, 350 ans\"."
                            },
                            {
                                "type": "list_item",
                                "value": "Marsan, Jean-Claude (1990). Montreal in evolution . McGill-Queen's University Press. ISBN 978-0-7735-0798-2 .",
                                "links": [
                                    {
                                        "url": "https://archive.org/details/montrealinevolut00jean",
                                        "text": "Montreal in evolution"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/ISBN_(identifier)",
                                        "text": "ISBN"
                                    },
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Special:BookSources/978-0-7735-0798-2",
                                        "text": "978-0-7735-0798-2"
                                    }
                                ]
                            },
                            {
                                "type": "list_item",
                                "value": "Tomàs, Mariona. \"Exploring the metropolitan trap: the case of Montreal.\" International Journal of Urban and Regional Research (2012) 36#3 pp: 554–567. doi:10.1111/j.1468-2427.2011.01066.x .",
                                "links": [
                                    {
                                        "url": "https://en.wikipedia.org/wiki/Doi_(identifier)",
                                        "text": "doi"
                                    },
                                    {
                                        "url": "https://doi.org/10.1111%2Fj.1468-2427.2011.01066.x",
                                        "text": "10.1111/j.1468-2427.2011.01066.x"
                                    }
                                ]
                            },
                            {
                                "type": "list_item",
                                "value": "\"2006 Census of Canada\" . Statistics Canada. 2008. Archived from the original on October 10, 2008. Retrieved May 28, 2008 .",
                                "links": [
                                    {
                                        "url": "https://web.archive.org/web/20081010164347/http://www12.statcan.ca/english/census/index.cfm",
                                        "text": "\"2006 Census of Canada\""
                                    },
                                    {
                                        "url": "http://www12.statcan.ca/english/census/index.cfm",
                                        "text": "the original"
                                    }
                                ]
                            },
                            {
                                "type": "list_item",
                                "value": "\"Montreal\" . 2006 Census of Canada: Community Profiles . Statistics Canada. 2008. Archived from the original on December 2, 2008. Retrieved May 28, 2008 .",
                                "links": [
                                    {
                                        "url": "https://web.archive.org/web/20081202162740/http://www12.statcan.ca/english/census06/data/profiles/community/Details/Page.cfm?Lang=E&Geo1=CD&Code1=2466&Geo2=PR&Code2=24&Data=Count&SearchText=Montreal&SearchType=Begins&SearchPR=01&B1=All&Custom=",
                                        "text": "\"Montreal\""
                                    },
                                    {
                                        "url": "http://www12.statcan.ca/english/census06/data/profiles/community/Details/Page.cfm?Lang=E&Geo1=CD&Code1=2466&Geo2=PR&Code2=24&Data=Count&SearchText=Montreal&SearchType=Begins&SearchPR=01&B1=All&Custom=",
                                        "text": "the original"
                                    }
                                ]
                            },
                            {
                                "type": "list_item",
                                "value": "Natural Resources Canada (2005). Canadian Geographical Names: Island of Montreal . Retrieved August 29, 2005.",
                                "links": [
                                    {
                                        "url": "https://web.archive.org/web/20080531042123/http://geonames.nrcan.gc.ca/education/montreal_e.php",
                                        "text": "Canadian Geographical Names: Island of Montreal"
                                    }
                                ]
                            },
                            {
                                "type": "list_item",
                                "value": "Michael Sletcher, \"Montréal\", in James Ciment, ed., Colonial America: An Encyclopedia of Social, Political, Cultural, and Economic History (5 vols., N.Y., 2005)."
                            }
                        ]
                    }
                ]
            }
        ]
    },
    .
    .
]
```
</detail>


ii) Get pre-parsed articles with name `Montreal` from English Wikipedia with selected fields. Example (here)[https://github.com/wikimedia-enterprise/wme-sdk-go/tree/feature/add-examples/example/structured-contents/example.md]