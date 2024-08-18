package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

type Question struct {
    QuestionText string   `json:"questionText"`
    ImageUrl     string   `json:"imageUrl"`
    Answer       []Answer `json:"answer"`
    Explanation  string   `json:"explanation"`
}

type Answer struct {
    Text   string `json:"text"`
    Answer bool   `json:"answer"`
}

var questions []Question

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next(w, r)
    }
}

func getQuestions(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(questions)
}

func main() {
    // Initialize questions
    questions = []Question{
        {
            QuestionText: "Is this a seven wonders of the world?",
            ImageUrl:     "assets/images/Taj_Mahal_(Edited).jpg",
            Answer: []Answer{
                {Text: "Yes", Answer: true},
                {Text: "No", Answer: false},
            },
            Explanation: "Yes, this is one of the seven wonders of the world",
        },
		{
            QuestionText: "Buy or Sell?",
            ImageUrl:     "https://thumbs.dreamstime.com/b/deciding-whether-to-buy-sell-bright-gold-sit-opposite-ends-gray-board-which-balanced-red-question-mark-40328461.jpg",
            Answer: []Answer{
                {Text: "Buy", Answer: true},
                {Text: "Sell", Answer: false},
            },
			Explanation: "With symmetrical triangles, you never know for sure which way it will break.That's why they're symmetrical - because there is an equilibrium between the buyers and sellers.However, the closer the price gets to the triangle's apex, the higher the chance of a strong breakout when one side takes over",
        },
        {
            QuestionText: "Is this the flag of India?",
            ImageUrl: "https://upload.wikimedia.org/wikipedia/commons/thumb/d/d9/Flag_of_Canada_%28Pantone%29.svg/1200px-Flag_of_Canada_%28Pantone%29.svg.png",
            Answer: []Answer{
                {Text: "Yes", Answer: false},
                {Text: "No", Answer: true},
            },
            Explanation: "This is the flag of Canada",
        },
        {
            QuestionText: "Is this the national bird of India?",
            ImageUrl: "https://i0.wp.com/doodlewash.com/wp-content/uploads/2018/11/Day-3-Peacock-Watercolor-Illustration-Patterns.jpg?fit=1024%2C692&ssl=1",
            Answer: []Answer{
                {Text: "Yes", Answer: true},
                {Text: "No", Answer: false},
            },
            Explanation: "Yes, this is the national bird of India",
        },
        {
            QuestionText: "Which IPL team is this?",
            ImageUrl: "https://i0.wp.com/www.smartprix.com/bytes/wp-content/uploads/2023/03/cover.png?ssl=1&quality=80&w=f",
            Answer: []Answer{
                {Text: "RCB", Answer: true},
                {Text: "CSK", Answer: false},
            }, 
            Explanation: "This is the logo of Royal Challengers Bangalore",
        },
        // Add more questions...
    }

    // Initialize router
    r := mux.NewRouter()

	r.HandleFunc("/questions", enableCORS(getQuestions)).Methods("GET", "OPTIONS")

    // Start server
    log.Println("Server starting on :8000")
    log.Fatal(http.ListenAndServe(":8000", r))
}
