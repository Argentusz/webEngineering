module.exports = {
  root: true,
  env: {
    node: true
  },
  extends: ["plugin:vue/essential"],
  rules: {
    "space-before-function-paren": ["warn", "never"],
    "key-spacing": [1, { "mode": "minimum" }],
    "max-depth": ["warn", 4],
    "no-multi-spaces": ["warn", { exceptions: { "VariableDeclarator": true } }],
    "quotes": [2, "double", { avoidEscape: true }],
    "prefer-const": ["warn", {
      "destructuring": "any",
      "ignoreReadBeforeAssign": false
    }]
  },
  parserOptions: {
    parser: "@babel/eslint-parser"
  }
}
