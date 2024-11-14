import { Breadcrumb, BreadcrumbItem } from 'react-bootstrap';
import { useTranslation } from 'react-i18next'

const SiteBreadCrumbs = (props) => {
  const { t, i18n } = useTranslation()
  const route = [{_id: '', name: t('menu.home')}, ...props.route]

  return (
    <Breadcrumb>
      { 
        route.map((site, idx) => {
          // Last element should be marked as active without a link
          const attributes = idx === route.length-1 
            ? {active: true}
            : {href: `/grower/${site._id}`,
               className: 'text-success'
            }

          return (
            <BreadcrumbItem
              key={idx}
              id={site._id}
              {...attributes}>
                {site.name}
            </BreadcrumbItem>
          )
        }
      )}
    </Breadcrumb>
  )
}

export default SiteBreadCrumbs
