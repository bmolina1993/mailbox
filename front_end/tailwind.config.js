/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  darkMode: "class",
  theme: {
    extend: {
      width: {
        4.5: ["1.125rem"] /* 18px */,
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
        menu: "rgba(0, 0, 0, 0.5)",
      }),
    },
    fontFamily: {
      roboto: ["Roboto", "sans-serif", "Segoe UI", "Ubuntu"],
    },
  },
  plugins: [require("tailwind-scrollbar")({ nocompatible: true })],
  variants: {
    scrollbar: ["rounded"],
  },
};
