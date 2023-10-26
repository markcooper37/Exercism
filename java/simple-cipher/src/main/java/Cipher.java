public class Cipher {

    private String key;

    public Cipher() {
        this.key = "aaaaaaaaaa";
    }

    public Cipher(String key) {
        this.key = key;
    }

    public String getKey() {
        return this.key;
    }

    public String encode(String plainText) {
        String encoding = "";
        for (int i = 0; i < plainText.length(); i++) {
            int encodingValue = this.key.charAt(i % this.key.length()) - 97;
            encoding += (char) (97 + ((plainText.charAt(i) - 97 + encodingValue) % 26));
        }
        return encoding;
    }

    public String decode(String cipherText) {
        String decoding = "";
        for (int i = 0; i < cipherText.length(); i++) {
            int encodingValue = this.key.charAt(i % this.key.length()) - 97;
            decoding += (char) (97 + ((cipherText.charAt(i) - 71 - encodingValue) % 26));
        }
        return decoding;
    }
}
