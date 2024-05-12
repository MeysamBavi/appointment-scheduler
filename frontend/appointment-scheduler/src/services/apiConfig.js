export default class apiConfig {
  static host = "http://135.125.116.88:80";

  static baseUrl() {
    return `${this.host}/api`;
  }

  static authBaseUrl() {
    return `${this.baseUrl()}/auth`;
  }

  static otpBaseUrl() {
    return `${this.authBaseUrl()}/otp`;
  }

  static otpSendUrl() {
    return `${this.otpBaseUrl()}/send`;
  }

  static otpValidateUrl() {
    return `${this.otpBaseUrl()}/validate`;
  }
}
