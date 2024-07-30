// icons
import {
  DashboardOutlined,
  IdcardOutlined,
  DatabaseOutlined,
  LineChartOutlined,
  MessageOutlined,
  CalendarOutlined,
  BuildOutlined,
  CustomerServiceOutlined,
  MailOutlined,
  ShoppingCartOutlined,
  UserOutlined,
  FileTextOutlined,
  PhoneOutlined,
  GoldOutlined,
  CloudUploadOutlined,
  FormOutlined,
  TableOutlined,
  PieChartOutlined,
  FileDoneOutlined,
  LoginOutlined,
  DollarOutlined,
  RocketOutlined,
  BellOutlined,
  QuestionOutlined,
  LockOutlined,
  CrownOutlined,
  MenuUnfoldOutlined,
  StopOutlined,
  BorderOutlined,
  BoxPlotOutlined,
  ChromeOutlined,
  DeploymentUnitOutlined
} from '@ant-design/icons-vue';

export interface menu {
  header?: string;
  title?: string;
  icon?: object;
  to?: string;
  getURL?: boolean;
  divider?: boolean;
  chip?: string;
  chipColor?: string;
  chipVariant?: string;
  chipIcon?: string;
  children?: menu[];
  disabled?: boolean;
  type?: string;
  subCaption?: string;
}

const sidebarItem: menu[] = [
  { header: 'Dashboard' },
  {
    title: 'Dashboard',
    icon: DashboardOutlined,
    to: '/dashboard/default',
    children: [
      {
        title: 'Default',
        to: '/dashboard/default'
      },
      {
        title: 'Analytics',
        to: '/dashboard/analytics'
      }
    ]
  },
  {
    title: 'Components',
    icon: GoldOutlined,
    to: 'components/buttons',
    getURL: true,
    type: 'external',
    chip: 'new',
    chipColor: 'primary',
    chipVariant: 'tonal'
  },
  { header: 'Widget' },
  {
    title: 'Statistics',
    icon: IdcardOutlined,
    to: '/widget/statistics'
  },
  {
    title: 'Data',
    icon: DatabaseOutlined,
    to: '/widget/data'
  },
  {
    title: 'Chart',
    icon: LineChartOutlined,
    to: '/widget/chart'
  },
  { header: 'Applications' },
  {
    title: 'Chat',
    icon: MessageOutlined,
    to: '/app/chats'
  },
  {
    title: 'Calendar',
    icon: CalendarOutlined,
    to: '/app/calendar'
  },
  {
    title: 'Kanban',
    icon: BuildOutlined,
    to: '/app/kanban'
  },
  {
    title: 'Customer',
    icon: CustomerServiceOutlined,
    to: '/customer/',
    children: [
      {
        title: 'Customer List',
        to: '/customer/customerlist'
      },
      {
        title: 'Create Invoice',
        to: '/app/customer/create-invoice'
      },
      {
        title: 'Order Details',
        to: '/app/customer/order-details'
      },
      {
        title: 'Order List',
        to: '/customer/orderlist'
      },
      {
        title: 'Product List',
        to: '/customer/productlist'
      },
      {
        title: 'Product Review',
        to: '/customer/productreview'
      }
    ]
  },
  {
    title: 'Invoice',
    icon: FileTextOutlined,
    to: '/',
    children: [
      {
        title: 'Dashboard',
        to: '/app/invoice/dashboard'
      },
      {
        title: 'Create',
        to: '/app/invoice/create'
      },
      {
        title: 'Details',
        to: '/app/invoice/details'
      },
      {
        title: 'List',
        to: '/app/invoice/list'
      },
      {
        title: 'Edit',
        to: '/app/invoice/edit'
      }
    ]
  },
  {
    title: 'Users',
    icon: UserOutlined,
    to: '/app/user',
    children: [
      {
        title: 'Social Profile',
        to: '/app/user/social/posts'
      },
      {
        title: 'Account Profile',
        to: '/app/user/account-profile',
        children: [
          {
            title: 'Profile 01',
            to: '/app/user/account-profile/profile1'
          },
          {
            title: 'Profile 02',
            to: '/app/user/account-profile/profile2'
          },
          {
            title: 'Profile 03',
            to: '/app/user/account-profile/profile3'
          }
        ]
      },
      {
        title: 'User Profile',
        to: '/app/user/userprofile'
      },
      {
        title: 'Cards',
        to: '/app/user/card',
        children: [
          {
            title: 'Style 01',
            to: '/app/user/card/card1'
          },
          {
            title: 'Style 02',
            to: '/app/user/card/card2'
          },
          {
            title: 'Style 03',
            to: '/app/user/card/card3'
          }
        ]
      },
      {
        title: 'List',
        to: '/app/user/list',
        children: [
          {
            title: 'Style 01',
            to: '/app/user/list1'
          },
          {
            title: 'Style 02',
            to: '/app/user/list2'
          }
        ]
      }
    ]
  },
  {
    title: 'Mail',
    icon: MailOutlined,
    to: '/app/mail'
  },
  {
    title: 'Contact',
    icon: PhoneOutlined,
    to: '/app/contacts',
    children: [
      {
        title: 'Card',
        to: '/app/contact/c-card'
      },
      {
        title: 'List',
        to: '/app/contact/c-list'
      }
    ]
  },
  {
    title: 'E-Commerce',
    icon: ShoppingCartOutlined,
    to: '/ecommerce/',
    children: [
      {
        title: 'Products',
        to: '/ecommerce/products'
      },
      {
        title: 'Product Detail',
        to: '/ecommerce/product/detail/1'
      },
      {
        title: 'Product List',
        to: '/ecommerce/productlist'
      },
      {
        title: 'Add New Product',
        to: '/ecommerce/add-product'
      },
      {
        title: 'Checkout',
        to: '/ecommerce/checkout'
      }
    ]
  },
  { header: 'Forms' },
  {
    title: 'Plugins',
    icon: CloudUploadOutlined,
    to: '/forms/radio',
    children: [
      {
        title: 'Editor',
        to: '/forms/plugins/editor'
      },
      {
        title: 'Mask',
        to: '/forms/plugins/mask'
      },
      {
        title: 'Captcha',
        to: '/auth/login1'
      },
      {
        title: 'Clipboard',
        to: '/forms/plugins/clipboard'
      }
    ]
  },

  {
    title: 'Layouts',
    icon: FormOutlined,
    to: '/forms/layouts',
    children: [
      {
        title: 'Layouts',
        to: '/forms/layouts/layouts'
      },
      {
        title: 'Multi Columns Form',
        to: '/forms/layouts/multi-column-forms'
      },
      {
        title: 'Action Bar',
        to: '/forms/layouts/action-bar'
      },
      {
        title: 'Sticky Action Bar',
        to: '/forms/layouts/sticky-action-bar'
      }
    ]
  },
  {
    title: 'Tables',
    icon: TableOutlined,
    to: '/forms/tables',
    children: [
      {
        title: 'Basic Table',
        to: '/tables/tbl-basic'
      },
      {
        title: 'Dark Table',
        to: '/tables/tbl-dark'
      },
      {
        title: 'Density Table',
        to: '/tables/tbl-density'
      },
      {
        title: 'Height Table',
        to: '/tables/tbl-height'
      },
      {
        title: 'Fixed Header Table',
        to: '/tables/tbl-fixed-header'
      }
    ]
  },
  {
    title: 'Charts',
    icon: PieChartOutlined,
    to: '/forms/radio',
    children: [
      {
        title: 'Org Charts',
        to: '/forms/charts/orgchart'
      },
      {
        title: 'Apex Charts',
        to: '/forms/charts/apexchart'
      }
    ]
  },
  {
    title: 'Form Validation',
    icon: FileDoneOutlined,
    to: '/forms/formvalidation'
  },
  { header: 'Utilities' },
  {
    title: 'Icons',
    icon: CrownOutlined,
    to: '/forms/radio',
    children: [
      {
        title: 'Ant design Icons',
        to: '/icons/ant'
      },
      {
        title: 'Tabler Icons',
        to: '/icons/tabler'
      },
      {
        title: 'Material Icons',
        to: '/icons/material'
      }
    ]
  },
  { divider: true },
  { header: 'Pages' },
  {
    title: 'Authentication',
    icon: LoginOutlined,
    to: '/auth',
    children: [
      {
        title: 'Login',
        to: '/auth/login1'
      },
      {
        title: 'Register',
        to: '/auth/register1'
      },
      {
        title: 'Forgot Password',
        to: '/auth/forgot-pwd1'
      },
      {
        title: 'Check Mail',
        to: '/auth/check-mail1'
      },
      {
        title: 'Reset Password',
        to: '/auth/reset-pwd1'
      },
      {
        title: 'Code Verification',
        to: '/auth/code-verify1'
      }
    ]
  },
  {
    title: 'Pricing',
    icon: DollarOutlined,
    to: '/pages/pricing1'
  },
  {
    title: 'Maintenance',
    icon: RocketOutlined,
    to: '/maintenenace',
    children: [
      {
        title: 'Error 404',
        to: '/pages/error'
      },
      {
        title: 'Error 500',
        to: '/pages/error500'
      },
      {
        title: 'Coming soon',
        to: '/pages/comingsoon1'
      },
      {
        title: 'Under Construction',
        to: '/pages/construction'
      }
    ]
  },
  {
    title: 'Landing page',
    icon: BellOutlined,
    getURL: true,
    to: '',
    type: 'external'
  },
  {
    title: 'Contact Us',
    icon: PhoneOutlined,
    getURL: true,
    to: 'pages/contact-us',
    type: 'external'
  },
  {
    title: 'FAQs',
    icon: QuestionOutlined,
    getURL: true,
    to: 'pages/faq',
    type: 'external'
  },
  {
    title: 'Privacy Policy',
    icon: LockOutlined,
    getURL: true,
    to: 'pages/privacy-policy',
    type: 'external'
  },
  { divider: true },
  { header: 'Others' },
  {
    title: 'Menu levels',
    icon: MenuUnfoldOutlined,
    to: '#',
    children: [
      {
        title: 'Level 1',
        to: ''
      },
      {
        title: 'Level 1',
        to: '',
        children: [
          {
            title: 'Level 2',
            to: ''
          },
          {
            title: 'Level 2',
            to: '/2.2level',
            children: [
              {
                title: 'Level 3',
                to: ''
              }
            ]
          }
        ]
      }
    ]
  },
  {
    title: 'Sub Caption',
    icon: BoxPlotOutlined,
    subCaption: 'caption title',
    to: ''
  },
  {
    title: 'Disabled Menu',
    icon: StopOutlined,
    disabled: true,
    to: ''
  },
  {
    title: 'Oval Chip',
    icon: BorderOutlined,
    to: ''
  },
  {
    title: 'Sample Page',
    icon: ChromeOutlined,
    to: '/starter'
  },
  {
    title: 'Documentation',
    icon: QuestionOutlined,
    to: 'https://codedthemes.gitbook.io/mantis-vuetify/',
    type: 'external',
    chip: 'gitbook',
    chipColor: 'secondary',
    chipVariant: 'flat'
  },
  {
    title: 'Road Map',
    icon: DeploymentUnitOutlined,
    to: 'https://codedthemes.gitbook.io/mantis-vuetify/roadmap',
    type: 'external'
  }
];

export default sidebarItem;
