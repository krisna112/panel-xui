[English](/README.md) | [ÙØ§Ø±Ø³ÛŒ](/README.fa_IR.md) | [Ø§Ù„Ø¹Ø±Ø¨ÙŠØ©](/README.ar_EG.md) |  [ä¸­æ–‡](/README.zh_CN.md) | [EspaÃ±ol](/README.es_ES.md) | [Ð ÑƒÑÑÐºÐ¸Ð¹](/README.ru_RU.md)

<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="./media/3x-ui-dark.png">
    <img alt="3x-ui" src="./media/3x-ui-light.png">
  </picture>
</p>

[![Release](https://img.shields.io/github/v/release/krisna112/panel-xui.svg)](https://github.com/krisna112/panel-xui/releases)
[![Build](https://img.shields.io/github/actions/workflow/status/krisna112/panel-xui/release.yml.svg)](https://github.com/krisna112/panel-xui/actions)
[![GO Version](https://img.shields.io/github/go-mod/go-version/krisna112/panel-xui.svg)](#)
[![Downloads](https://img.shields.io/github/downloads/krisna112/panel-xui/total.svg)](https://github.com/krisna112/panel-xui/releases/latest)
[![License](https://img.shields.io/badge/license-GPL%20V3-blue.svg?longCache=true)](https://www.gnu.org/licenses/gpl-3.0.en.html)
[![Go Reference](https://pkg.go.dev/badge/github.com/krisna112/panel-xui/v2.svg)](https://pkg.go.dev/github.com/krisna112/panel-xui/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/krisna112/panel-xui/v2)](https://goreportcard.com/report/github.com/krisna112/panel-xui/v2)

**3X-UI** â€” Ð¿Ñ€Ð¾Ð´Ð²Ð¸Ð½ÑƒÑ‚Ð°Ñ Ð¿Ð°Ð½ÐµÐ»ÑŒ ÑƒÐ¿Ñ€Ð°Ð²Ð»ÐµÐ½Ð¸Ñ Ñ Ð¾Ñ‚ÐºÑ€Ñ‹Ñ‚Ñ‹Ð¼ Ð¸ÑÑ…Ð¾Ð´Ð½Ñ‹Ð¼ ÐºÐ¾Ð´Ð¾Ð¼ Ð½Ð° Ð¾ÑÐ½Ð¾Ð²Ðµ Ð²ÐµÐ±-Ð¸Ð½Ñ‚ÐµÑ€Ñ„ÐµÐ¹ÑÐ°, Ñ€Ð°Ð·Ñ€Ð°Ð±Ð¾Ñ‚Ð°Ð½Ð½Ð°Ñ Ð´Ð»Ñ ÑƒÐ¿Ñ€Ð°Ð²Ð»ÐµÐ½Ð¸Ñ ÑÐµÑ€Ð²ÐµÑ€Ð¾Ð¼ Xray-core. ÐŸÑ€ÐµÐ´Ð¾ÑÑ‚Ð°Ð²Ð»ÑÐµÑ‚ ÑƒÐ´Ð¾Ð±Ð½Ñ‹Ð¹ Ð¸Ð½Ñ‚ÐµÑ€Ñ„ÐµÐ¹Ñ Ð´Ð»Ñ Ð½Ð°ÑÑ‚Ñ€Ð¾Ð¹ÐºÐ¸ Ð¸ Ð¼Ð¾Ð½Ð¸Ñ‚Ð¾Ñ€Ð¸Ð½Ð³Ð° Ñ€Ð°Ð·Ð»Ð¸Ñ‡Ð½Ñ‹Ñ… VPN Ð¸ Ð¿Ñ€Ð¾ÐºÑÐ¸-Ð¿Ñ€Ð¾Ñ‚Ð¾ÐºÐ¾Ð»Ð¾Ð².

> [!IMPORTANT]
> Ð­Ñ‚Ð¾Ñ‚ Ð¿Ñ€Ð¾ÐµÐºÑ‚ Ð¿Ñ€ÐµÐ´Ð½Ð°Ð·Ð½Ð°Ñ‡ÐµÐ½ Ñ‚Ð¾Ð»ÑŒÐºÐ¾ Ð´Ð»Ñ Ð»Ð¸Ñ‡Ð½Ð¾Ð³Ð¾ Ð¸ÑÐ¿Ð¾Ð»ÑŒÐ·Ð¾Ð²Ð°Ð½Ð¸Ñ, Ð¿Ð¾Ð¶Ð°Ð»ÑƒÐ¹ÑÑ‚Ð°, Ð½Ðµ Ð¸ÑÐ¿Ð¾Ð»ÑŒÐ·ÑƒÐ¹Ñ‚Ðµ ÐµÐ³Ð¾ Ð² Ð½ÐµÐ·Ð°ÐºÐ¾Ð½Ð½Ñ‹Ñ… Ñ†ÐµÐ»ÑÑ… Ð¸ Ð² Ð¿Ñ€Ð¾Ð¸Ð·Ð²Ð¾Ð´ÑÑ‚Ð²ÐµÐ½Ð½Ð¾Ð¹ ÑÑ€ÐµÐ´Ðµ.

ÐšÐ°Ðº ÑƒÐ»ÑƒÑ‡ÑˆÐµÐ½Ð½Ð°Ñ Ð²ÐµÑ€ÑÐ¸Ñ Ð¾Ñ€Ð¸Ð³Ð¸Ð½Ð°Ð»ÑŒÐ½Ð¾Ð³Ð¾ Ð¿Ñ€Ð¾ÐµÐºÑ‚Ð° X-UI, 3X-UI Ð¾Ð±ÐµÑÐ¿ÐµÑ‡Ð¸Ð²Ð°ÐµÑ‚ Ð¿Ð¾Ð²Ñ‹ÑˆÐµÐ½Ð½ÑƒÑŽ ÑÑ‚Ð°Ð±Ð¸Ð»ÑŒÐ½Ð¾ÑÑ‚ÑŒ, Ð±Ð¾Ð»ÐµÐµ ÑˆÐ¸Ñ€Ð¾ÐºÑƒÑŽ Ð¿Ð¾Ð´Ð´ÐµÑ€Ð¶ÐºÑƒ Ð¿Ñ€Ð¾Ñ‚Ð¾ÐºÐ¾Ð»Ð¾Ð² Ð¸ Ð´Ð¾Ð¿Ð¾Ð»Ð½Ð¸Ñ‚ÐµÐ»ÑŒÐ½Ñ‹Ðµ Ñ„ÑƒÐ½ÐºÑ†Ð¸Ð¸.

## Ð‘Ñ‹ÑÑ‚Ñ€Ñ‹Ð¹ ÑÑ‚Ð°Ñ€Ñ‚

```
bash <(curl -Ls https://raw.githubusercontent.com/krisna112/panel-xui/main/install.sh)
```

ÐŸÐ¾Ð»Ð½ÑƒÑŽ Ð´Ð¾ÐºÑƒÐ¼ÐµÐ½Ñ‚Ð°Ñ†Ð¸ÑŽ ÑÐ¼Ð¾Ñ‚Ñ€Ð¸Ñ‚Ðµ Ð² [Ð²Ð¸ÐºÐ¸ Ð¿Ñ€Ð¾ÐµÐºÑ‚Ð°](https://github.com/krisna112/panel-xui/wiki).

## ÐžÑÐ¾Ð±Ð°Ñ Ð±Ð»Ð°Ð³Ð¾Ð´Ð°Ñ€Ð½Ð¾ÑÑ‚ÑŒ

- [alireza0](https://github.com/alireza0/)

## Ð‘Ð»Ð°Ð³Ð¾Ð´Ð°Ñ€Ð½Ð¾ÑÑ‚Ð¸

- [Iran v2ray rules](https://github.com/chocolate4u/Iran-v2ray-rules) (Ð›Ð¸Ñ†ÐµÐ½Ð·Ð¸Ñ: **GPL-3.0**): _Ð£Ð»ÑƒÑ‡ÑˆÐµÐ½Ð½Ñ‹Ðµ Ð¿Ñ€Ð°Ð²Ð¸Ð»Ð° Ð¼Ð°Ñ€ÑˆÑ€ÑƒÑ‚Ð¸Ð·Ð°Ñ†Ð¸Ð¸ Ð´Ð»Ñ v2ray/xray Ð¸ v2ray/xray-clients ÑÐ¾ Ð²ÑÑ‚Ñ€Ð¾ÐµÐ½Ð½Ñ‹Ð¼Ð¸ Ð¸Ñ€Ð°Ð½ÑÐºÐ¸Ð¼Ð¸ Ð´Ð¾Ð¼ÐµÐ½Ð°Ð¼Ð¸ Ð¸ Ñ„Ð¾ÐºÑƒÑÐ¾Ð¼ Ð½Ð° Ð±ÐµÐ·Ð¾Ð¿Ð°ÑÐ½Ð¾ÑÑ‚ÑŒ Ð¸ Ð±Ð»Ð¾ÐºÐ¸Ñ€Ð¾Ð²ÐºÑƒ Ñ€ÐµÐºÐ»Ð°Ð¼Ñ‹._
- [Russia v2ray rules](https://github.com/runetfreedom/russia-v2ray-rules-dat) (Ð›Ð¸Ñ†ÐµÐ½Ð·Ð¸Ñ: **GPL-3.0**): _Ð­Ñ‚Ð¾Ñ‚ Ñ€ÐµÐ¿Ð¾Ð·Ð¸Ñ‚Ð¾Ñ€Ð¸Ð¹ ÑÐ¾Ð´ÐµÑ€Ð¶Ð¸Ñ‚ Ð°Ð²Ñ‚Ð¾Ð¼Ð°Ñ‚Ð¸Ñ‡ÐµÑÐºÐ¸ Ð¾Ð±Ð½Ð¾Ð²Ð»ÑÐµÐ¼Ñ‹Ðµ Ð¿Ñ€Ð°Ð²Ð¸Ð»Ð° Ð¼Ð°Ñ€ÑˆÑ€ÑƒÑ‚Ð¸Ð·Ð°Ñ†Ð¸Ð¸ V2Ray Ð½Ð° Ð¾ÑÐ½Ð¾Ð²Ðµ Ð´Ð°Ð½Ð½Ñ‹Ñ… Ð¾ Ð·Ð°Ð±Ð»Ð¾ÐºÐ¸Ñ€Ð¾Ð²Ð°Ð½Ð½Ñ‹Ñ… Ð´Ð¾Ð¼ÐµÐ½Ð°Ñ… Ð¸ Ð°Ð´Ñ€ÐµÑÐ°Ñ… Ð² Ð Ð¾ÑÑÐ¸Ð¸._

## ÐŸÐ¾Ð´Ð´ÐµÑ€Ð¶ÐºÐ° Ð¿Ñ€Ð¾ÐµÐºÑ‚Ð°

**Ð•ÑÐ»Ð¸ ÑÑ‚Ð¾Ñ‚ Ð¿Ñ€Ð¾ÐµÐºÑ‚ Ð¿Ð¾Ð»ÐµÐ·ÐµÐ½ Ð´Ð»Ñ Ð²Ð°Ñ, Ð²Ñ‹ Ð¼Ð¾Ð¶ÐµÑ‚Ðµ Ð¿Ð¾ÑÑ‚Ð°Ð²Ð¸Ñ‚ÑŒ ÐµÐ¼Ñƒ**:star2:

<a href="https://www.buymeacoffee.com/MHSanaei" target="_blank">
<img src="./media/default-yellow.png" alt="Buy Me A Coffee" style="height: 70px !important;width: 277px !important;" >
</a>

</br>
<a href="https://nowpayments.io/donation/hsanaei" target="_blank" rel="noreferrer noopener">
   <img src="./media/donation-button-black.svg" alt="Crypto donation button by NOWPayments">
</a>

## Ð—Ð²ÐµÐ·Ð´Ñ‹ Ñ Ñ‚ÐµÑ‡ÐµÐ½Ð¸ÐµÐ¼ Ð²Ñ€ÐµÐ¼ÐµÐ½Ð¸

[![Stargazers over time](https://starchart.cc/krisna112/panel-xui.svg?variant=adaptive)](https://starchart.cc/krisna112/panel-xui) 
