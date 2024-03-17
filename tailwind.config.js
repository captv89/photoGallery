/** @type {import('tailwindcss').Config} */

const defaultTheme = require('tailwindcss/defaultTheme');


module.exports = {
	content: ['./web/tfs/**/*.{templ,html,js,go}'],
	theme: {
		extend: {
			fontFamily: {
				sans: ['"Proxima Nova"', ...defaultTheme.fontFamily.sans],
			},
		},
	},
	plugins: [],
};

