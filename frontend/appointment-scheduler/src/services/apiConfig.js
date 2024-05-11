

// TODO: create a hierarchial architecture for reuse baseUrl
module.exports = {
    // --------------------------------- bases
    hostname: 'http://localhost/api',
    port: 32768,
    baseUrl: 'http://localhost/api:23768',

    // --------------------------------- full urls
    otpSendUrl: 'http://localhost/api:23768/otp/send',
    otpValidateUrl: 'http://localhost/api:23768/otp/validate'
}