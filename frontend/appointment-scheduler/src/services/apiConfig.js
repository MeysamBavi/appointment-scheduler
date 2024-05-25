export default class apiConfig {
  static host = "http://135.125.116.88:80";
  // static host = "http://localhost:8080";

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

  static businessUrl() {
    return `${this.baseUrl()}/business-manager`;
  }

  static userUrl() {
    return `${this.businessUrl()}/businesses`;
  }

  static businessTypeUrl() {
    return `${this.businessUrl()}/service_types`;
  }
}
