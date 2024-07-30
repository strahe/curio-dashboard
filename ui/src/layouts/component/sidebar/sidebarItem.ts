export interface menu {
  header?: string;
  title?: string;
  to?: string;
}

const sidebarItem: menu[] = [
  { header: 'Inputs' },
  {
    title: 'Autocomplete',
    to: '/components/autocomplete'
  },
  {
    title: 'Buttons',
    to: '/components/buttons'
  },

  {
    title: 'Checkbox',
    to: '/components/checkbox'
  },
  {
    title: 'Radio',
    to: '/components/radio'
  },
  {
    title: 'Ratings',
    to: '/advance/ratings'
  },
  {
    title: 'Switch',
    to: '/components/switch'
  },
  {
    title: 'Slider',
    to: '/components/slider'
  },
  {
    title: 'TextField',
    to: '/components/textfield'
  },
  { header: 'Data Display' },
  {
    title: 'Avatar',
    to: '/basic/avatar'
  },
  {
    title: 'Badges',
    to: '/basic/badges'
  },
  {
    title: 'Chip',
    to: '/basic/chip'
  },
  {
    title: 'List',
    to: '/basic/uilist'
  },
  {
    title: 'Tooltip',
    to: '/forms/plugins/tooltip'
  },
  {
    title: 'Typography',
    to: '/utils/typography'
  },
  { header: 'Feedback' },
  {
    title: 'Alert',
    to: '/advance/alert'
  },
  {
    title: 'Dialog',
    to: '/forms/plugins/modal'
  },
  {
    title: 'Progress',
    to: '/advance/progress'
  },
  {
    title: 'Snackbar',
    to: '/advance/snackbar'
  },
  { header: 'Navigation' },
  {
    title: 'Breadcrumb',
    to: '/basic/breadcrumb'
  },
  {
    title: 'Pagination',
    to: '/advance/pagination'
  },
  {
    title: 'Tabs',
    to: '/basic/tabs'
  },
  { header: 'Surfaces' },
  {
    title: 'Accordion',
    to: '/basic/expansion-panel'
  },
  {
    title: 'Cards',
    to: '/basic/cards'
  },
  { header: 'Utils' },
  {
    title: 'Colors',
    to: '/utils/colors'
  },
  {
    title: 'Shadows',
    to: '/utils/shadows'
  },
  {
    title: 'Timeline',
    to: '/advance/timeline'
  },
  {
    title: 'Button Toggle',
    to: '/advance/btn-toggle'
  },
  {
    title: 'Grid',
    to: '/utils/grid'
  },
  {
    title: 'Animation',
    to: '/utils/animation'
  }
];

export default sidebarItem;
