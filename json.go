package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string, logErr error) {
	if logErr != nil {
		log.Println(logErr)
	}
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	/*//w.Write(dat)
	//bu yazma işlemi sırasında hata olursa, hatayı logluyoruz
	//ama kullanıcıya zaten başlık gönderildiği için, hatayı iletemiyoruz
	//bu yüzden sadece logluyoruz
	//eğer hata yoksa, veriyi yazıyoruz
	//bu işlem başarılı olursa, kullanıcıya veri gönderilmiş oluyor
	//eğer hata olursa, kullanıcıya veri gönderilememiş oluyor
	// veriyi nerede döndürüyoruz? w.Write(dat) ile nasıl veriyi döndürüyoruz? ama return yok?
	//w.Write(dat) fonksiyonu, veriyi HTTP yanıtına yazar
	//eğer bu işlem başarılı olursa, veri kullanıcıya gönderilmiş olur
	//eğer bu işlem sırasında hata olursa, hatayı logluyoruz
	//ama fonksiyonun sonunda return yok
	//çünkü fonksiyonun sonunda return yazmaya gerek yok
	//çünkü fonksiyonun sonunda zaten işlem bitiyor nasıl biitiyor?
	//fonksiyonun sonunda işlem bittiğinde, fonksiyonun çağrıldığı yere geri dönülüyor
	//bu yüzden fonksiyonun sonunda return yazmaya gerek yok
	//ama fonksiyonun ortasında return yazmak gerekiyor
	//çünkü fonksiyonun ortasında işlem bitiyor ve fonksiyonun çağrıldığı yere geri dönülüyor
	//örneğin json.Marshal(payload) işlemi sırasında hata olursa, hatayı logluyoruz ve fonksiyondan çıkıyoruz
	//çünkü bu durumda veriyi yazamıyoruz
	// hata aldığı için return var yani
	//ama en azından hatayı loglamış oluyoruz
	//böylece hatanın ne olduğunu anlayabiliriz
	//ve gerektiğinde müdahale edebiliriz
	//_,err := w.Write(dat) neden _ kullanıyoruz?
	//çünkü w.Write fonksiyonu iki değer döner
	//birincisi yazılan byte sayısı
	//ikincisi hata durumu
	//ama biz sadece hata durumunu kontrol etmek istiyoruz
	//yazılan byte sayısını kullanmıyoruz
	//bu yüzden _ kullanıyoruz
	//_ kullanmazsak, Go derleyicisi hata verir
	//çünkü kullanılmayan değişkenler Go dilinde hata olarak kabul edilir
	//bu yüzden _ kullanarak bu hatayı önlüyoruz
	//eğer sadece w.Write(dat) olarak çağırırsak, derleyici hata verir
	//çünkü dönen değerlerden biri kullanılmıyor
	//bu yüzden _ kullanmak zorundayız
	//önceden sadece w.Write(dat) yazıyorduk bu yüzden hata alıyorduk çünkü dönen değerlerden biri kullanılmıyordu
	//w.Write(dat) tek başına çalıştırılamaz çünkü dönen değerlerden biri kullanılmıyor
	*/
	_, err = w.Write(dat)
	if err != nil {
		log.Printf("Error writing response: %s", err)
	}
}
