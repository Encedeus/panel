/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{html,js,svelte,ts}"],
  theme: {
    extend: {
      fontFamily: {
        "inter": "Inter",
        "monti": "Montserrat",
        "lato": "Lato",
      }
    },
  },
  plugins: [],
}

