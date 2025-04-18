# Changelog

Version numbers use the following system:

    0.X.Y,

where first 0 indicates pre-release version, X tells which Software engineering project Slowers group has done the release and Y is used to tell which sprint the release was made in.

## Release 0.2.6
[GitHub release](https://github.com/Slowers-Team/Slowers-App/releases/tag/v0.2.6)

- Backend refactored to use Postgres for user, business and business memberships and Mongo for sites, flowers and images
- Removed USESQL feature toggle
- Deprecated old staging server
- Business owners can see information about their business and employees in Business page
- Business owner can add/edit/remove other users as employees in their business

## Release 0.2.5
[GitHub release](https://github.com/Slowers-Team/Slowers-App/releases/tag/v0.2.5)

- Images are stored in Cloudinary
- User accounts and business information are saved in Postgres-database
- Frontend uses containers for page contents

Still using two staging servers, one with working routing and Postgres-database enabled (uses [Dockerfile in the root](../Dockerfile)) and one with working site, flower and image handling (uses separate Dockerfiles for [backend](../backend/Dockerfile) and [frontend](../frontend/Dockerfile)).

## Release 0.2.4
[GitHub release](https://github.com/Slowers-Team/Slowers-App/releases/tag/v0.2.4)

- Routing in the staging server is fixed
    - Two staging servers, one with working routing and one with working addition of images
- Background image is added
- Business creation form has new fields
- Flowers are hidden after pre-determined length of time
- General fixes and preparations to move user, business and business membership storing to Postgres-database

## Release 0.2.3
[GitHub release](https://github.com/Slowers-Team/Slowers-App/releases/tag/0.2.3)

- Logo for the app
- Thumbnail creation
- Business creation format and Business homepage, There are two options Retailer Business or Grower Business
- User can no longer change between grower and retailer from the profile

## Changes made in versions 0.2.1-0.2.2
Features done by the second group working on the project in sprints 1 and 2.

- Marketplace has filters for scientific name, grower email and flower name
- Removed unnecessary pop-ups
- Business owner home page
- Swedish language support
- Bug fixes

## Features in version 0.1
Features added by the first group working on the project.

- User can register and log in as grower or retailer
- Growers can create fields
- Growers can add flowers
- Growers can choose to put flowers to the Marketplace
- Retailers can see flowers added by Growers in the Marketplace
- Language support for English and Finnish
