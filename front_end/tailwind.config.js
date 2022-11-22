/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  darkMode: "class",
  theme: {
    extend: {
      backgroundImage: {
        // sanFrancisco: "url('/public/img/sanFrancisco.jpg')",
      },

      colors: {
        darkPrimary: "#1E293B",
        darkSecondary: "#293548",
        primary: "#7986CC",
        secondary: "#E06055",
        tertiary: "#4CB6AC",
        fourth: "#9CCC66",
      },

      backgroundColor: (theme) => ({
        ...theme("colors"),
        darkPrimary: "#1E293B",
        darkSecondary: "#293548",
        primary: "#7986CC",
        secondary: "#E06055",
        tertiary: "#4CB6AC",
        fourth: "#9CCC66",
      }),

      /*
      textColor: {
        primary: "#CC2D4A",
        secondary: "#8FA206",
        tertiary: "#61AEC9",
      },
      */
    },
    fontFamily: {
      roboto: ["Roboto", "sans-serif", "Segoe UI", "Ubuntu"],
    },
  },
  /**/
  plugins: [require("tailwind-scrollbar")({ nocompatible: true })],
  variants: {
    scrollbar: ["rounded"],
  },
};
