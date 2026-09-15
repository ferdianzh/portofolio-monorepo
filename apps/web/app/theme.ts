import { createTheme } from "@mantine/core";

export const theme = createTheme({
  fontFamily: "Geist, sans-serif",
  headings: {
    fontFamily: "Fraunces, serif",
  },

  colors: {
    primary: [
      "#e5fdf8", // 1
      "#d7f4ee", // 2
      "#b4e5dc", // 3
      "#8ed6c8", // 4
      "#6ec9b7", // 5
      "#58c1ad", // 6
      "#44bba4", // 7
      "#39a792", // 8
      "#2c9581", // 9
      "#11816f", // 10
    ],
    secondary: [
      "#fff8e2", // 1
      "#fbf0cf", // 2
      "#f5dfa4", // 3
      "#eece74", // 4
      "#e7bb41", // 5
      "#e5b632", // 6
      "#e4b122", // 7
      "#ca9b13", // 8
      "#b48908", // 9
      "#9c7600", // 10
    ],
    dark: [
      "#f4f5f5", // 1
      "#e7e7e7", // 2
      "#cdcdcd", // 3
      "#b2b2b2", // 4
      "#9b9b9b", // 5
      "#8c8c8c", // 6
      "#848585", // 7
      "#717273", // 8
      "#636667", // 9
      "#393e41", // 10
    ],
    light: [
      "#f8f6ee", // 1
      "#e7e5df", // 2
      "#d4d1c9", // 3
      "#bcb7ab", // 4
      "#a7a191", // 5
      "#9b9480", // 6
      "#958d76", // 7
      "#817a64", // 8
      "#736c56", // 9
      "#645d46", // 10
    ],
  },

  primaryColor: "primary",
  primaryShade: 5,
});
