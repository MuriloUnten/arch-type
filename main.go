package main

import(
    "fmt";
    "log";
    "os";
    "encoding/json";
    "io";
    "strconv";
    "math/rand/v2"
)

type Mode string
const MODE_WORDS Mode = "words"
const MODE_TIME Mode = "time"


type Settings struct {
    Mode Mode
    Time int
    Words int
    Language Language
}


type Language struct {
    Name string `json:"name"`
    Words []string `json:"words"`
    NumWords int
}


func (settings *Settings) SetLanguage(langStr string) {
    langPath := "languages/" + langStr + ".json"
    jsonFile, err := os.Open(langPath)
    if err != nil {
        log.Fatal(err)
    }
    defer jsonFile.Close()

    jsonByteValue, _ := io.ReadAll(jsonFile)
    json.Unmarshal(jsonByteValue, &(settings.Language))
    settings.Language.NumWords = len(settings.Language.Words)
}


func GenerateTest(settings *Settings) []string {
    var wordSlice []string

    if (settings.Mode == MODE_WORDS) {
        wordSlice = make([]string, settings.Words)

        for i := 0; i < settings.Words; i++ {
            wordIdx := rand.IntN(settings.Language.NumWords)
            wordSlice[i] = settings.Language.Words[wordIdx]
        }
    } else if (settings.Mode == MODE_TIME) {
        numOfPredefWords := 25
        wordSlice = make([]string, numOfPredefWords)

        for i := 0; i < numOfPredefWords; i++ {
            wordIdx := rand.IntN(settings.Language.NumWords)
            wordSlice[i] = settings.Language.Words[wordIdx]
        }
    }

    return wordSlice
}


func main() {
    var settings Settings
    settings.Mode = MODE_TIME
    settings.Time = 15
    settings.Words = 10
    settings.SetLanguage("english")

    fmt.Println("Mode: " + settings.Mode)
    fmt.Println("Time: " + strconv.Itoa(settings.Time))
    fmt.Println("Words: " + strconv.Itoa(settings.Words))
    fmt.Println("Language: " + settings.Language.Name)

    test := GenerateTest(&settings)
    fmt.Println(test)
}
