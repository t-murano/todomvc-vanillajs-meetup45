module.exports = function(config) {
  config.set({
    frameworks: ['jasmine'],
    files: [
      '../js/controller.js',
      'ControllerSpec.js'
    ],
    customLaunchers: {
      Chrome_without_security: {
        base: 'Chrome',
        flags: ['--disable-web-security']
      }
    }
  })
}
