import CryptoJS from 'crypto-js';

const key = CryptoJS.enc.Utf8.parse('aimzc886'); // 密钥，后端提供
// const iv = CryptoJS.enc.Utf8.parse(''); // 偏移量，这里为空，通常在CBC模式下使用

/**
 * AES加密：字符串 key iv 返回base64
 * @param word 要加密的字符串
 * @returns 加密后的base64字符串
 */
export function Encrypt(word: string): string {
  const srcs = CryptoJS.enc.Utf8.parse(word);
  const encrypted = CryptoJS.AES.encrypt(srcs, key, {
    iv:CryptoJS.enc.Utf8.parse(''),
    mode: CryptoJS.mode.ECB,
    padding: CryptoJS.pad.Pkcs7,
  });
  return CryptoJS.enc.Base64.stringify(encrypted.ciphertext);
}

/**
 * AES解密：字符串 key iv 返回base64
 * @param word 要解密的base64编码字符串
 * @returns 解密后的字符串
 */
export function Decrypt(word: string): string {
  const base64 = CryptoJS.enc.Base64.parse(word);
  const src = CryptoJS.enc.Base64.stringify(base64);

  const decrypt = CryptoJS.AES.decrypt(src, key, {
    iv:CryptoJS.enc.Utf8.parse(''),
    mode: CryptoJS.mode.ECB,
    padding: CryptoJS.pad.Pkcs7,
  });
  return CryptoJS.enc.Utf8.stringify(decrypt);
}