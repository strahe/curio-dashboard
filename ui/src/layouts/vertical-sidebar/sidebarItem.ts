// icons
import {
  CrownOutlined,
  CustomerServiceOutlined,
  DashboardOutlined,
  HomeOutlined,
  IdcardOutlined,
  QuestionOutlined,
  TableOutlined,
} from '@ant-design/icons-vue'

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
    title: 'Overview',
    icon: HomeOutlined,
    to: '/dashboard/overview',
  },
  {
    title: 'Analytics',
    icon: DashboardOutlined,
    to: '/dashboard/default',
  },
  { header: 'Cluster' },
  {
    title: 'Machines',
    icon: IdcardOutlined,
    to: '/app/machines',
  },
  {
    title: 'Storages',
    icon: IdcardOutlined,
    to: '/app/storages',
  },
  { header: 'Sector' },
  {
    title: 'Customer',
    icon: CustomerServiceOutlined,
    to: '/customer/',
    children: [
      {
        title: 'Customer List',
        to: '/customer/customerlist',
      },
      {
        title: 'Create Invoice',
        to: '/app/customer/create-invoice',
      },
      {
        title: 'Order Details',
        to: '/app/customer/order-details',
      },
      {
        title: 'Order List',
        to: '/customer/orderlist',
      },
      {
        title: 'Product List',
        to: '/customer/productlist',
      },
      {
        title: 'Product Review',
        to: '/customer/productreview',
      },
    ],
  },
  { header: 'Forms' },
  {
    title: 'Tables',
    icon: TableOutlined,
    to: '/forms/tables',
    children: [
      {
        title: 'Basic Table',
        to: '/tables/tbl-basic',
      },
      {
        title: 'Dark Table',
        to: '/tables/tbl-dark',
      },
      {
        title: 'Density Table',
        to: '/tables/tbl-density',
      },
      {
        title: 'Height Table',
        to: '/tables/tbl-height',
      },
      {
        title: 'Fixed Header Table',
        to: '/tables/tbl-fixed-header',
      },
    ],
  },
  { header: 'Utilities' },
  {
    title: 'Icons',
    icon: CrownOutlined,
    to: '/forms/radio',
    children: [
      {
        title: 'Ant design Icons',
        to: '/icons/ant',
      },
      {
        title: 'Tabler Icons',
        to: '/icons/tabler',
      },
      {
        title: 'Material Icons',
        to: '/icons/material',
      },
    ],
  },
  { header: 'Pages' },
  {
    title: 'FAQs',
    icon: QuestionOutlined,
    getURL: true,
    to: 'pages/faq',
    type: 'external',
  },
  { divider: true },
]

export default sidebarItem
