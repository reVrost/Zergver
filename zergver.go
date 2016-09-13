package main

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
	//	"net/smtp"
	"strconv"
)

func home(writer http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(writer, "Hello... darkness my old friend")
	fmt.Println("Endpoint Hit: Main")
}

func returnMailSetting(writer http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(writer, "Hello... this is mail setting")
	fmt.Println("Endpoint Hit: Mail Setting")
}

func getMailRecipient(writer http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(writer, "Hello... this is get mail recipient")
	fmt.Println("Endpoint Hit: Set Mail Recipient")
}

func getMailMessage(writer http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(writer, "Hello... this is get mail message")
	fmt.Println("Endpoint Hit: Set Mail Message")

}

const DB_NAME string = "zergdb"

type zergver struct {
	db *bolt.DB
}

func initZergver(dbfile string) (z *zergver, err error) {
	z = &zergver{}
	z.db, err = bolt.Open(dbfile, 0600, &bolt.Options{Timeout: 1 * time.Second})

	z.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(DB_NAME))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	return
}

func saveRecipientHandle(z *zergver) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		//		decoder := json.NewDecoder(r.Body)
		//		var emails []string
		//		err := decoder.Decode(&emails)
		//		b, err := tx.CreateBucketIfNotExists([]byte("posts"))
		//		if err != nil {
		//			return err
		//		}
		encoded, err := json.Marshal(r.Body)

		if err != nil {
			panic(err)
		}

		z.db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(DB_NAME))

			err := b.Put([]byte("recipient"), []byte(encoded))
			return err

		})
		fmt.Fprintf(w, "Hello... this is get mail message %s", encoded)
	}
}

func getRecipientHandle(z *zergver) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		err := z.db.View(func(tx *bolt.Tx) error {

			b := tx.Bucket([]byte(DB_NAME))
			v := b.Get([]byte("recipient"))
			fmt.Println("The recipients are: %s\n", v)
			return nil
		})

		if err != nil {
			log.Fatalf("Errorget lul: %s", err)
		}
	}
}

func setMailRecipient(writer http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(writer, "Hello... this is set mail recipient")
	fmt.Println("Endpoint Hit: Set Mail Recipient")
}

func setMailMessage(writer http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(writer, "Hello... this is set mail message")
	fmt.Println("Endpoint Hit: Set Mail Message")

}

func startServer(port string) {
	router := httprouter.New()

	z, err := initZergver("zergver.db")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	router.GET("/", home)
	router.GET("/mail-setting", returnMailSetting)
	router.GET("/mail-recipient", setMailRecipient)

	router.GET("/get-recipient", getRecipientHandle(z))
	router.POST("/save-recipient", saveRecipientHandle(z))

	log.Fatal(http.ListenAndServe(":"+port, router))
}

func main() {
	port := strconv.Itoa(4242)
	fmt.Println("Starting zergver at " + port)
	startServer(port)
}
