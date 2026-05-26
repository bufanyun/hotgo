import CryptoJS from 'crypto-js';

const defaultKey = 'f080a463654b2279';

export const aesEcb = {
  // 加密
  encrypt(word: string, keyStr: string = defaultKey): string {
    const key = CryptoJS.enc.Utf8.parse(keyStr);
    const src = CryptoJS.enc.Utf8.parse(word);
    const encrypted = CryptoJS.AES.encrypt(src, key, {
      mode: CryptoJS.mode.ECB,
      padding: CryptoJS.pad.Pkcs7,
    });
    return encrypted.toString();
  },
  // 解密
  decrypt(word: string, keyStr: string = defaultKey): string {
    const key = CryptoJS.enc.Utf8.parse(keyStr);
    const decrypt = CryptoJS.AES.decrypt(word, key, {
      mode: CryptoJS.mode.ECB,
      padding: CryptoJS.pad.Pkcs7,
    });
    return CryptoJS.enc.Utf8.stringify(decrypt).toString();
  },
};
