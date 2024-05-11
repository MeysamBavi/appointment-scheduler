

export default class apiConfig {
    static hostname = import.meta.env.VITE_API_HOSTNAME;
    static port = import.meta.env.VITE_API_PORT;

    static baseUrl() {
        return `${this.hostname}:${this.port}/api`
    }

    static otpBaseUrl() {
        return `${this.baseUrl()}/otp`
    }

    static otpSendUrl() {
        return `${this.otpBaseUrl()}/send`
    }

    static otpValidateUrl() {
        return `${this.otpBaseUrl()}/validate`
    }
}