package auth

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

// *testing.T: test fonksiyonlarının çalıştırılması ve sonuçların raporlanması için kullanılır
// burada testing.T neden pointer olarak kullanılıyor? çünkü test fonksiyonları genellikle büyük veri yapılarıyla çalışabilir ve bu verilerin kopyalanması performans sorunlarına yol açabilir
// bunun dışında kopyalanmadığı için test fonksiyonları içinde yapılan değişiklikler orijinal test nesnesine yansır ve bu da test sonuçlarının doğru bir şekilde raporlanmasını sağlar
// eğer *testing.T yerine testing.T kullanılsaydı, test fonksiyonları içinde yapılan değişiklikler orijinal test nesnesine yansımayacaktı ve bu da test sonuçlarının doğru bir şekilde raporlanmamasına neden olacaktı
// örneğin fonksiyon içinde t.Fail() çağrılırsa, bu değişiklik orijinal test nesnesine yansımayacak ve test başarısız olarak raporlanmayacaktı
// *testing paketi, Go dilinde test yazmak için kullanılan standart bir pakettir
// struct içinde tanımlanan kullanılmak zorunda olduğumuz değişkenler
// key: HTTP başlığının anahtarı (örneğin, "Authorization")
// value: HTTP başlığının değeri (örneğin, "ApiKey xxxxx")
// expect: beklenen API anahtarı sonucu
// expectErr: beklenen hata mesajı
// boş bırakılan alanlar için varsayılan değerler kullanılır
// örneğin, key boşsa Authorization başlığı eklenmez bunun yerine boş bırakılır ve bu durumda GetAPIKey fonksiyonu hata döner
// peki yanlış değerler girilirse ne olur? örneğin value alanı boş bırakılırsa Authorization başlığı eklenir ancak değeri boş olur ve bu durumda GetAPIKey fonksiyonu yine hata döner
// bu şekilde farklı senaryolar test edilebilir
// her test durumu için beklenen sonuçlar ve hata mesajları belirtilir
// böylece fonksiyonun doğru çalışıp çalışmadığı kontrol edilir
// value içinde verilen - işareti, Authorization başlığının değeri olarak kullanılır
// bu durumda GetAPIKey fonksiyonu "malformed authorization header" hatası döner
// value içinde verilen "Bearer xxxxx" değeri, Authorization başlığının değeri olarak kullanılır
// bu durumda GetAPIKey fonksiyonu yine "malformed authorization header" hatası döner çünkü beklenen format "ApiKey xxxxx" şeklindedir
// value içinde verilen "ApiKey xxxxx" değeri, Authorization başlığının değeri olarak kullanılır
// bu durumda GetAPIKey fonksiyonu "invalid api key" hatası döner çünkü "xxxxx" geçerli bir API anahtarı değildir
// value içinde verilen "ApiKey valid_api_key" değeri, Authorization başlığının değeri olarak kullanılır
// bu durumda GetAPIKey fonksiyonu başarılı olur ve "valid_api_key" sonucunu döner
func TestApiKeyAuthCases(t *testing.T) {
	tests := []struct {
		key       string
		value     string
		expect    string
		expectErr string
	}{
		{
			expectErr: "no authorization header",
		},
		{
			key:       "Authorization",
			expectErr: "no authorization header",
		},
		{
			key:       "Authorization",
			value:     "-",
			expectErr: "malformed authorization header",
		},
		{
			key:       "Authorization",
			value:     "Bearer xxxxx",
			expectErr: "malformed authorization header",
		},
		{
			key:       "Authorization",
			value:     "ApiKey ",
			expectErr: "invalid api key",
		},
		{
			key:       "Authorization",
			value:     "ApiKey valid_api_key",
			expect:    "valid_api_key",
			expectErr: "not expected to be an error",
		},
	}
	/* bu satırda yukarıdaki testleri dolaşıyoruz i index test içerik*/
	for i, test := range tests {
		// Her test için alt test çalıştırıyoruz
		// fmt.Sprintf ile test adını oluşturuyoruz
		// t.Run ile alt testi başlatıyoruz
		// http.Header yapısını oluşturup test verilerini ekliyoruz
		// GetAPIKey fonksiyonunu çağırıyoruz ve sonucu kontrol ediyoruz
		// Hata varsa beklenen hata mesajıyla karşılaştırıyoruz
		// Beklenen sonuçla karşılaştırma yapıyoruz
		// Hataları t.Errorf ile raporluyoruz
		// Testin başarılı olup olmadığını kontrol ediyoruz
		// Testin sonunda döngü devam ediyor ve bir sonraki testi işliyoruz
		// Bu şekilde tüm testler tek tek çalıştırılmış oluyor
		// Her testin sonucu bağımsız olarak değerlendiriliyor
		// Test çıktıları konsola yazdırılıyor
		// Testler tamamlandığında genel bir sonuç raporu oluşturulabilir
		// Bu rapor testlerin başarı durumunu özetler
		// Gerekirse ek hata ayıklama bilgileri de eklenebilir
		// Test süreci böylece tamamlanmış olur
		// t.Run fonksiyonu içinde testin kendisi tanımlanır
		// Bu yapı testlerin düzenli ve okunabilir olmasını sağlar
		// Her test kendi başına çalıştırılabilir ve izole edilebilir
		// Bu da hata ayıklamayı kolaylaştırır
		// Test fonksiyonları genellikle küçük ve odaklanmış olur
		// Bu da kodun bakımını kolaylaştırır
		// Testler genellikle hızlı çalışır, bu da geliştirme sürecini hızlandırır
		// t.Run içindeki fonksiyon anonim bir fonksiyondur
		// Bu fonksiyon testin kendisini temsil eder
		// Testin adı fmt.Sprintf ile dinamik olarak oluşturulur
		// Bu sayede her testin adı benzersiz olur
		// Test çıktıları daha anlaşılır hale gelir
		// bu anonim fonksiyon parametre olarak *testing.T alır
		// kontrol yapıları bu parametre üzerinden gerçekleştirilir
		// t.Errorf ile hata mesajları raporlanır
		t.Run(fmt.Sprintf("TestGetAPIKey Case #%v:", i), func(t *testing.T) {
			header := http.Header{}
			header.Add(test.key, test.value)

			output, err := GetAPIKey(header)
			if err != nil {
				if strings.Contains(err.Error(), test.expectErr) {
					return
				}
				t.Errorf("Unexpected: TestGetAPIKey:%v\n", err)
				return
			}

			if output != test.expect {
				t.Errorf("Unexpected: TestGetAPIKey:%s", output)
				return
			}
		})
	}
}
