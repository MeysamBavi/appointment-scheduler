
module.exports = class apiConfig {
    static hostname = 'http://localhost/api';
    static port = 32768;

    static baseUrl() {
        return `${this.hostname}:${this.port}`
    }

    static otpBaseUrl() {
        return `${this.baseUrl()}/otp`
    }

    static otpSendUrl() {
        return `${this.otpBaseUrl()}/send`
    }

    static otpValidateUrl() {
        return `${this.otpSendUrl()}/validate`
    }
}