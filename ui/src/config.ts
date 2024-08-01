export type ConfigProps = {
  Sidebar_drawer: boolean;
  Customizer_drawer: boolean;
  mini_sidebar: boolean;
  setHorizontalLayout: boolean;
  actTheme: string;
  fontTheme: string;
  inputBg: boolean;
  boxed: boolean;
  dark: boolean;
};

const config: ConfigProps = {
  Sidebar_drawer: true,
  Customizer_drawer: false,
  mini_sidebar: false,
  setHorizontalLayout: false, // Horizontal layout
  actTheme: 'DefaultTheme',
  fontTheme: 'Public sans',
  inputBg: false,
  boxed: false,
  dark: false,
}

export default config
