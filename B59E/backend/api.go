package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var (
	regexEmail   = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	regexDollars = regexp.MustCompile(`^[0-9]*$`)
	regexID      = regexp.MustCompile("^[a-zA-Z0-9]*$")
	regexName    = regexp.MustCompile("^[a-zA-Z ]*$")
	regexBirth   = regexp.MustCompile(`\d{2}/\d{2}/\d{4}`)
)

// Propose the launch project content
type Propose struct {
	Title       string `json:"title" bson:"title"`
	Name        string `json:"name" bson:"name"`
	Email       string `json:"email" bson:"email"`
	TargetFund  string `json:"targetFund" bson:"targetFund"`
	CurrentFund string `json:"currentFund" bson:"currentFund"`
	Enddate     string `json:"enddate" bson:"enddate"`
	Description string `json:"description" bson:"description"`
	URL         string `json:"url" bson:"url"`
}

type Register struct {
	ID             string `json:"id" bson:"id"`
	Password       string `json:"password" bson:"password"`
	Email          string `json:"email" bson:"email"`
	Name           string `json:"name" bson:"name"`
	Identification string `json:"identification" bson:"identification"`
	Birth          string `json:"birth" bson:"birth"`
	Registertime   string `json:"registertime" bson:"registertime"`
	Token          string `json:"token" bson:"token"`
}

// RegisterUser define user's detail

func handleLogin(c *gin.Context) {
	var parser struct {
		ID       string `json:"id" bson:"id"`
		Password string `json:"password" bson:"password"`
	}
	rawdata, err := c.GetRawData()
	if err != nil {
		log.Println("ERROR JSON Raw Data")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 2.Unserialize
	if err := json.Unmarshal(rawdata, &parser); err != nil {
		log.Println("ERROR Json Key")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 3. Regex Check
	if regexID.MatchString(parser.ID) &&
		len(parser.Password) >= 6 {

		// 4.Find Name in Mongodb
		res, err := FindOne("user", "account", &parser)
		if err != nil {
			log.Println("Error find the DB")
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		b, err := bson.Marshal(res)
		if err != nil {
			log.Println("Error in bson Marshal")
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err.Error()})
			return
		}
		parser := Register{}
		if err = bson.Unmarshal(b, &parser); err != nil {
			log.Println("Error Json Unmarshal")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"id":             parser.ID,
			"email":          parser.Email,
			"name":           parser.Name,
			"identification": parser.Identification,
			"birth":          parser.Birth,
			"token":          parser.Token,
		})
		return
	}
}

func handlePropose(c *gin.Context) {
	// 2. Parser the purpose info
	parser := Propose{}
	rawdata, err := c.GetRawData()
	if err != nil {
		log.Println("ERROR Json Raw Data")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(string(rawdata))
	// 2.Unserialize
	if err = json.Unmarshal(rawdata, &parser); err != nil {
		log.Println("ERROR Json Key")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 3. Regex Check
	if regexEmail.MatchString(parser.Email) &&
		regexDollars.MatchString(parser.TargetFund) {
		// 4. Insert To DB
		log.Println("Insert in MongoDB")
		res, err := InsertOne("user", "propose", &parser)
		if err != nil {
			log.Println("Insert Fail")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Println(res)
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	}
}

func handleGetPropose(c *gin.Context) {
	const databaseName = "user"
	const collectionName = "propose"
	// Connect to mongodb
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Error connect to Mongodb", err)
	}
	defer cancel()
	// Query Name
	collection := client.Database(databaseName).Collection(collectionName)
	cur, err := collection.Find(context.Background(), bson.D{})
	num, err := collection.CountDocuments(context.Background(), bson.D{})
	defer cur.Close(ctx)
	// 	If find, return false
	parserAll := make([]Propose, num)
	cur.All(ctx, &parserAll)
	log.Println(parserAll)

	c.JSON(http.StatusOK, parserAll)
}

func handleQueryName(c *gin.Context) {
	var parser struct {
		ID       string `json:"id" bson:"id"`
		Password string `json:"password" bson:"password"`
		Email    string `json:"email" bson:"email"`
	}
	// 1.Check Receive Data
	rawdata, err := c.GetRawData()
	if err != nil {
		log.Println("ERROR Json Raw Data")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 2.Unserialize
	if err = json.Unmarshal(rawdata, &parser); err != nil {
		log.Println("ERROR Json Key")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 3. Regex Check
	if regexID.MatchString(parser.ID) && regexEmail.MatchString(parser.Email) && len(parser.Password) >= 6 {
		// 4.Find Name in Mongodb
		log.Println("start query....")
		parse := bson.D{
			{"$or",
				bson.A{
					bson.D{{"id", parser.ID}},
					bson.D{{"email", parser.Email}}},
			}}
		res, err := Find("user", "account", parse)
		// 5. If find the email or id in db, then
		if err != nil {
			log.Println("Error Query in mongodb")
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		}
		if len(res) == 0 {
			c.JSON(http.StatusOK, gin.H{"status": "OK"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "ID Already in used"})
			return
		}
		return
	} else {
		log.Fatal("ERROR Json Key")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func handleRegister(c *gin.Context) {
	parser := Register{}
	// 1.Check Receive Data
	rawdata, err := c.GetRawData()
	if err != nil {
		log.Println("ERROR Json Raw Data")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 2.Unserialize
	err = json.Unmarshal(rawdata, &parser)
	if err != nil {
		log.Println("ERROR Json Key")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 3. Regex Check
	if regexID.MatchString(parser.ID) &&
		regexEmail.MatchString(parser.Email) &&
		regexName.MatchString(parser.Name) &&
		regexBirth.MatchString(parser.Birth) &&
		len(parser.Password) >= 6 {
		// 4. Generate JWT token
		token, err := GenerateToken(parser.ID, parser.Password)
		parser.Token = token
		if err != nil {
			log.Println("Error Generate JWT token")
		}
		log.Println("Insert in MongoDB")

		_, err = InsertOne("user", "account", &parser)
		if err != nil {
			log.Println("Insert Fail")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insert Fail"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"token":  parser.Token,
		})
	}
}

func handleAuth(c *gin.Context) {
	token := c.PostForm("token")
	if token == "" {
		log.Println("No token")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "bad token"})
		return
	}
	claims, err := ParseToken(token)
	if err != nil {
		log.Println("Error Auth check")
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	} else if time.Now().Unix() > claims.ExpiresAt {
		log.Println("Error token timeout")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token timeout"})
		return
	}
	query := bson.D{
		{"$and",
			bson.A{
				bson.D{{"id", claims.ID}},
				bson.D{{"token", token}}},
		}}
	res, err := FindOne("user", "account", &query)
	if err != nil {
		log.Println("Error find the DB")
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	b, err := bson.Marshal(res)
	if err != nil {
		log.Println("Error in bson Marshal")
		c.JSON(http.StatusExpectationFailed, gin.H{"error": err.Error()})
		return
	}
	parser := Register{}
	if err = bson.Unmarshal(b, &parser); err != nil {
		log.Println("Error Json Unmarshal")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":             parser.ID,
		"email":          parser.Email,
		"name":           parser.Name,
		"identification": parser.Identification,
		"Birth":          parser.Birth,
	})
	return
}
