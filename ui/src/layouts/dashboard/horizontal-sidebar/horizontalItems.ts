import { CircleFilledIcon } from 'vue-tabler-icons';
// icons
import {
  DashboardOutlined,
  DesktopOutlined,
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
  PhoneOutlined,
  GoldOutlined,
  CloudUploadOutlined,
  FormOutlined,
  TableOutlined,
  PieChartOutlined,
  FileDoneOutlined,
  BellOutlined,
  QuestionOutlined,
  CrownOutlined,
  AppstoreAddOutlined,
  StopOutlined,
  BorderOutlined,
  ChromeOutlined,
  DeploymentUnitOutlined
} from '@ant-design/icons-vue';

export interface menu {
  header?: string;
  title?: string;
  icon?: object;
  to?: string;
  divider?: boolean;
  chip?: string;
  chipColor?: string;
  chipVariant?: string;
  chipIcon?: string;
  children?: menu[];
  disabled?: boolean;
  subCaption?: string;
  class?: string;
  extraclass?: string;
  type?: string;
}

const horizontalItems: menu[] = [
  {
    title: 'Dashboard',
    icon: DashboardOutlined,
    to: '#',
    children: [
      {
        title: 'Dashboard',
        icon: DashboardOutlined,
        to: '#',
        children: [
          {
            title: 'Default',
            icon: DashboardOutlined,
            to: '/dashboard/default'
          },
          {
            title: 'Analytics',
            icon: DesktopOutlined,
            to: '/dashboard/analytics'
          }
        ]
      },
      {
        title: 'Components',
        icon: GoldOutlined,
        to: '/components/buttons'
      }
    ]
  },
  {
    title: 'Widgets',
    icon: IdcardOutlined,
    to: '#',
    children: [
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
      }
    ]
  },
  {
    title: 'Apps',
    icon: AppstoreAddOutlined,
    to: '#',
    children: [
      {
        title: 'Users',
        icon: UserOutlined,
        to: '/app/user',
        children: [
          {
            title: 'Social Profile',
            icon: CircleFilledIcon,
            to: '/app/user/social/posts'
          },
          {
            title: 'Account Profile',
            icon: CircleFilledIcon,
            to: '/app/user/account-profile',
            children: [
              {
                title: 'Profile 01',
                icon: CircleFilledIcon,
                to: '/app/user/account-profile/profile1'
              },
              {
                title: 'Profile 02',
                icon: CircleFilledIcon,
                to: '/app/user/account-profile/profile2'
              },
              {
                title: 'Profile 03',
                icon: CircleFilledIcon,
                to: '/app/user/account-profile/profile3'
              }
            ]
          },
          {
            title: 'User Profile',
            icon: CircleFilledIcon,
            to: '/app/user/userprofile'
          },
          {
            title: 'Cards',
            icon: CircleFilledIcon,
            to: '/app/user/card',
            children: [
              {
                title: 'Style 01',
                icon: CircleFilledIcon,
                to: '/app/user/card/card1'
              },
              {
                title: 'Style 02',
                icon: CircleFilledIcon,
                to: '/app/user/card/card2'
              },
              {
                title: 'Style 03',
                icon: CircleFilledIcon,
                to: '/app/user/card/card3'
              }
            ]
          },
          {
            title: 'List',
            icon: CircleFilledIcon,
            to: '/app/user/list',
            children: [
              {
                title: 'Style 01',
                icon: CircleFilledIcon,
                to: '/app/user/list1'
              },
              {
                title: 'Style 02',
                icon: CircleFilledIcon,
                to: '/app/user/list2'
              }
            ]
          }
        ]
      },
      {
        title: 'Customer',
        icon: CustomerServiceOutlined,
        to: '/customer/',
        children: [
          {
            title: 'Customer List',
            icon: CircleFilledIcon,
            to: '/customer/customerlist'
          },
          {
            title: 'Create Invoice',
            icon: CircleFilledIcon,
            to: '/app/customer/create-invoice'
          },
          {
            title: 'Order Details',
            icon: CircleFilledIcon,
            to: '/app/customer/order-details'
          },
          {
            title: 'Order List',
            icon: CircleFilledIcon,
            to: '/customer/orderlist'
          },
          {
            title: 'Product List',
            icon: CircleFilledIcon,
            to: '/customer/productlist'
          },
          {
            title: 'Product Review',
            icon: CircleFilledIcon,
            to: '/customer/productreview'
          }
        ]
      },
      {
        title: 'Chat',
        icon: MessageOutlined,
        to: '/app/chats'
      },
      {
        title: 'Kanban',
        icon: BuildOutlined,
        to: '/app/kanban'
      },
      {
        title: 'Mail',
        icon: MailOutlined,
        to: '/app/mail'
      },
      {
        title: 'Calendar',
        icon: CalendarOutlined,
        to: '/app/calendar'
      },
      {
        title: 'Contact',
        icon: PhoneOutlined,
        to: '/app/contacts',
        children: [
          {
            title: 'Card',
            icon: CircleFilledIcon,
            to: '/app/contact/c-card'
          },
          {
            title: 'List',
            icon: CircleFilledIcon,
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
            icon: CircleFilledIcon,
            to: '/ecommerce/products'
          },
          {
            title: 'Product Detail',
            icon: CircleFilledIcon,
            to: '/ecommerce/product/detail/1'
          },
          {
            title: 'Product List',
            icon: CircleFilledIcon,
            to: '/ecommerce/productlist'
          },
          {
            title: 'Add New Product',
            icon: CircleFilledIcon,
            to: '/ecommerce/add-product'
          },
          {
            title: 'Checkout',
            icon: CircleFilledIcon,
            to: '/ecommerce/checkout'
          }
        ]
      }
    ]
  },
  {
    title: 'Forms',
    icon: FileDoneOutlined,
    to: '#',
    children: [
      {
        title: 'Plugins',
        icon: CloudUploadOutlined,
        to: '/forms/radio',
        children: [
          {
            title: 'Editor',
            icon: CircleFilledIcon,
            to: '/forms/plugins/editor'
          },
          {
            title: 'Mask',
            icon: CircleFilledIcon,
            to: '/forms/plugins/mask'
          },
          {
            title: 'Captcha',
            icon: CircleFilledIcon,
            to: '/auth/login1'
          },
          {
            title: 'Clipboard',
            icon: CircleFilledIcon,
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
            icon: CircleFilledIcon,
            to: '/forms/layouts/layouts'
          },
          {
            title: 'Multi Columns Form',
            icon: CircleFilledIcon,
            to: '/forms/layouts/multi-column-forms'
          },
          {
            title: 'Action Bar',
            icon: CircleFilledIcon,
            to: '/forms/layouts/action-bar'
          },
          {
            title: 'Sticky Action Bar',
            icon: CircleFilledIcon,
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
            icon: CircleFilledIcon,
            to: '/tables/tbl-basic'
          },
          {
            title: 'Dark Table',
            icon: CircleFilledIcon,
            to: '/tables/tbl-dark'
          },
          {
            title: 'Density Table',
            icon: CircleFilledIcon,
            to: '/tables/tbl-density'
          },
          {
            title: 'Height Table',
            icon: CircleFilledIcon,
            to: '/tables/tbl-height'
          },
          {
            title: 'Fixed Header Table',
            icon: CircleFilledIcon,
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
            icon: CircleFilledIcon,
            to: '/forms/charts/orgchart'
          },
          {
            title: 'Apex Charts',
            icon: CircleFilledIcon,
            to: '/forms/charts/apexchart'
          }
        ]
      },
      {
        title: 'Form Validation',
        icon: FileDoneOutlined,
        to: '/forms/formvalidation'
      }
    ]
  },
  {
    title: 'More items',
    to: '#',
    children: [
      {
        title: 'Front-end',
        icon: BellOutlined,
        to: '/2level',
        children: [
          {
            title: 'Landingpage',
            icon: CircleFilledIcon,
            to: '/',
            type: 'external'
          },
          {
            title: 'Contact Us',
            icon: CircleFilledIcon,
            to: '/pages/contact-us',
            type: 'external'
          },
          {
            title: 'FAQs',
            icon: CircleFilledIcon,
            to: '/pages/faq',
            type: 'external'
          },
          {
            title: 'Privacy Policy',
            icon: CircleFilledIcon,
            to: '/pages/privacy-policy',
            type: 'external'
          }
        ]
      },
      {
        title: 'Utilities',
        icon: CrownOutlined,
        to: '/level1',
        children: [
          {
            title: 'Ant design Icons',
            icon: CircleFilledIcon,
            to: '/icons/ant'
          },
          {
            title: 'Tabler Icons',
            icon: CircleFilledIcon,
            to: '/icons/tabler'
          },
          {
            title: 'Material Icons',
            icon: CircleFilledIcon,
            to: '/icons/material'
          }
        ]
      },
      {
        title: 'Disabled Menu',
        icon: StopOutlined,
        disabled: true,
        to: '/mantis'
      },
      {
        title: 'Oval Chip',
        icon: BorderOutlined,
        to: '/mantis'
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
        type: 'external'
      },
      {
        title: 'Road Map',
        icon: DeploymentUnitOutlined,
        to: 'https://codedthemes.gitbook.io/mantis-vuetify/roadmap',
        type: 'external'
      }
    ]
  }
];

export default horizontalItems;
