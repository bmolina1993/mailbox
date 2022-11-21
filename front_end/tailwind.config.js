/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  darkMode: "class",
  theme: {
    extend: {
      backgroundImage: {
        // sanFrancisco: "url('/public/img/sanFrancisco.jpg')",
      },
      backgroundColor: (theme) => ({
        ...theme("colors"),
        primary: "#CC2D4A",
        secondary: "#8FA206",
        tertiary: "#61AEC9",
      }),
      textColor: {
        primary: "#CC2D4A",
        secondary: "#8FA206",
        tertiary: "#61AEC9",
      },
    },
  },
  /*
  plugins: [require("tailwind-scrollbar")],
  variants: {
    scrollbar: ["rounded"],
  },
  */
};
