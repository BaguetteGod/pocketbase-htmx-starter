/** @type {import('tailwindcss').Config} */
module.exports = {
    mode: "jit",
    content: ["./**/*.html", "./**/*.templ", "./**/*.go"],
    plugins: [require("@tailwindcss/forms")],
    theme: {
        extend: {},
    },
    corePlugins: {
        preflight: true,
    },
};
