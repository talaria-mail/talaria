<p align="center">
  <img src="./docs/assets/talaria.svg" alt="Talaria" width="250px"/>
</p>

# Talaria

[![Build Status][ci-badge]][ci-link]
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Ftalaria-mail%2Ftalaria.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Ftalaria-mail%2Ftalaria?ref=badge_shield)

Talaria is an effort to create an email server that goes out of its way to make
it easy for you to host your own email.

**Goals**

- Low resource usage (1 vCPU, 500MiB Ram should be able to comfortably run Talaria)
- Easy configuration (Don't rely on docs to get users to set up tricky DNS, ask for AWS creds (or equivalent) and go set it up for them)
- Target only modern protocols (implicit TLS on submission, implicit TLS on IMAP, no support of POP at all etc)
- Stay away from email black lists with rigorous compliance to DKIM, SPF and other identity protocols (without making the user think about this stuff!)

**Non-Goals**

- Exhaustive compliance with all protocols
- High scalability or high availability

[ci-badge]: https://circleci.com/gh/talaria-mail/talaria.svg?style=shield
[ci-link]: https://circleci.com/gh/talaria-mail/talaria


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Ftalaria-mail%2Ftalaria.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Ftalaria-mail%2Ftalaria?ref=badge_large)