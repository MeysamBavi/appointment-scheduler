export default class apiConfig {
  static host = "http://94.23.161.62:80";
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
  static getEmployees(id) {
    return `${this.baseUrl()}/business-manager/businesses/${id}/employees`;
  }
  static businessUrl() {
    return `${this.baseUrl()}/business-manager`;
  }

  static businessesListUrl() {
    return `${this.businessUrl()}/businesses`;
  }

  static businessTypeUrl() {
    return `${this.businessUrl()}/service_types`;
  }
}
