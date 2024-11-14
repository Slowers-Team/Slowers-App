import { BreadcrumbItem } from 'react-bootstrap';
import Breadcrumb from 'react-bootstrap/Breadcrumb';


const SiteBreadCrumbs = ({ route }) => {
  return (
    <>
      <Breadcrumb>
        { 
          route.map((site, idx) =>
            <BreadcrumbItem
              key={idx}
              href={`/grower/${site._id}`}
              active={idx === route.length-1}>
                {site.name}
            </BreadcrumbItem>
    )}
      </Breadcrumb>
    </>
  )
}

export default SiteBreadCrumbs
