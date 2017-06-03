package emma

// Emma.css 0.4.0
const Src string = `
    ( pos-s       , position               , static ),
    ( pos-a       , position               , absolute ),
    ( pos-r       , position               , relative ),
    ( pos-f       , position               , fixed ),
    ( t-a         , top                    , auto ),
    ( t-0         , top                    , 0 ),
    ( r-a         , right                  , auto ),
    ( r-0         , right                  , 0 ),
    ( b-a         , bottom                 , auto ),
    ( b-0         , bottom                 , 0 ),
    ( l-a         , left                   , auto ),
    ( l-0         , left                   , 0 ),
    ( z-a         , z-index                , auto ),
    ( z-0         , z-index                , 0 ),
    ( z-1         , z-index                , 1 ),
    ( z-2         , z-index                , 2 ),
    ( z-3         , z-index                , 3 ),
    ( z-4         , z-index                , 4 ),
    ( z-5         , z-index                , 5 ),
    ( z-6         , z-index                , 6 ),
    ( fl-n        , float                  , none ),
    ( fl-l        , float                  , left ),
    ( fl-r        , float                  , right ),
    ( cl-n        , clear                  , none ),
    ( cl-l        , clear                  , left ),
    ( cl-r        , clear                  , right ),
    ( cl-b        , clear                  , both ),
    ( d-n         , display                , none ),
    ( d-b         , display                , block ),
    ( d-f         , display                , flex ),
    ( d-if        , display                , inline-flex ),
    ( d-i         , display                , inline ),
    ( d-ib        , display                , inline-block ),
    ( d-li        , display                , list-item ),
    ( d-ri        , display                , run-in ),
    ( d-cp        , display                , compact ),
    ( d-tb        , display                , table ),
    ( d-itb       , display                , inline-table ),
    ( d-tbcp      , display                , table-caption ),
    ( d-tbcl      , display                , table-column ),
    ( d-tbclg     , display                , table-column-group ),
    ( d-tbhg      , display                , table-header-group ),
    ( d-tbfg      , display                , table-footer-group ),
    ( d-tbr       , display                , table-row ),
    ( d-tbrg      , display                , table-row-group ),
    ( d-tbc       , display                , table-cell ),
    ( d-rb        , display                , ruby ),
    ( d-rbb       , display                , ruby-base ),
    ( d-rbbg      , display                , ruby-base-group ),
    ( d-rbt       , display                , ruby-text ),
    ( d-rbtg      , display                , ruby-text-group ),
    ( v-v         , visibility             , visible ),
    ( v-h         , visibility             , hidden ),
    ( v-c         , visibility             , collapse ),
    ( ov-v        , overflow               , visible ),
    ( ov-h        , overflow               , hidden ),
    ( ov-s        , overflow               , scroll ),
    ( ov-a        , overflow               , auto ),
    ( ovx-v       , overflow-x             , visible ),
    ( ovx-h       , overflow-x             , hidden ),
    ( ovx-s       , overflow-x             , scroll ),
    ( ovx-a       , overflow-x             , auto ),
    ( ovy-v       , overflow-y             , visible ),
    ( ovy-h       , overflow-y             , hidden ),
    ( ovy-s       , overflow-y             , scroll ),
    ( ovy-a       , overflow-y             , auto ),
    ( bxz-cb      , box-sizing             , content-box ),
    ( bxz-bb      , box-sizing             , border-box ),
    ( bxsh-n      , box-shadow             , none ),
    ( m-a         , margin                 , auto ),
    ( m-0         , margin                 , 0 ),
    ( m-0_a       , margin                 , 0 auto ),              // ^1
    ( m-xs        , margin                 , $emma-margin-xs ),     // ^1
    ( m-sm        , margin                 , $emma-margin-sm ),     // ^1
    ( m-md        , margin                 , $emma-margin-md ),     // ^1
    ( m-lg        , margin                 , $emma-margin-lg ),     // ^1
    ( m-xl        , margin                 , $emma-margin-xl ),     // ^1
    ( mt-0        , margin-top             , 0 ),
    ( mt-xs       , margin-top             , $emma-margin-xs ),     // ^1
    ( mt-sm       , margin-top             , $emma-margin-sm ),     // ^1
    ( mt-md       , margin-top             , $emma-margin-md ),     // ^1
    ( mt-lg       , margin-top             , $emma-margin-lg ),     // ^1
    ( mt-xl       , margin-top             , $emma-margin-xl ),     // ^1
    ( mr-a        , margin-right           , auto ),
    ( mr-0        , margin-right           , 0 ),
    ( mr-xs       , margin-right           , $emma-margin-xs ),     // ^1
    ( mr-sm       , margin-right           , $emma-margin-sm ),     // ^1
    ( mr-md       , margin-right           , $emma-margin-md ),     // ^1
    ( mr-lg       , margin-right           , $emma-margin-lg ),     // ^1
    ( mr-xl       , margin-right           , $emma-margin-xl ),     // ^1
    ( mb-0        , margin-bottom          , 0 ),
    ( mb-xs       , margin-bottom          , $emma-margin-xs ),     // ^1
    ( mb-sm       , margin-bottom          , $emma-margin-sm ),     // ^1
    ( mb-md       , margin-bottom          , $emma-margin-md ),     // ^1
    ( mb-lg       , margin-bottom          , $emma-margin-lg ),     // ^1
    ( mb-xl       , margin-bottom          , $emma-margin-xl ),     // ^1
    ( ml-a        , margin-left            , auto ),
    ( ml-0        , margin-left            , 0 ),
    ( ml-xs       , margin-left            , $emma-margin-xs ),     // ^1
    ( ml-sm       , margin-left            , $emma-margin-sm ),     // ^1
    ( ml-md       , margin-left            , $emma-margin-md ),     // ^1
    ( ml-lg       , margin-left            , $emma-margin-lg ),     // ^1
    ( ml-xl       , margin-left            , $emma-margin-xl ),     // ^1
    ( p-0         , padding                , 0 ),
    ( p-xs        , padding                , $emma-padding-xs ),    // ^1
    ( p-sm        , padding                , $emma-padding-sm ),    // ^1
    ( p-md        , padding                , $emma-padding-md ),    // ^1
    ( p-lg        , padding                , $emma-padding-lg ),    // ^1
    ( p-xl        , padding                , $emma-padding-xl ),    // ^1
    ( pt-0        , padding-top            , 0 ),
    ( pt-xs       , padding-top            , $emma-padding-xs ),    // ^1
    ( pt-sm       , padding-top            , $emma-padding-sm ),    // ^1
    ( pt-md       , padding-top            , $emma-padding-md ),    // ^1
    ( pt-lg       , padding-top            , $emma-padding-lg ),    // ^1
    ( pt-xl       , padding-top            , $emma-padding-xl ),    // ^1
    ( pr-0        , padding-right          , 0 ),
    ( pr-xs       , padding-right          , $emma-padding-xs ),    // ^1
    ( pr-sm       , padding-right          , $emma-padding-sm ),    // ^1
    ( pr-md       , padding-right          , $emma-padding-md ),    // ^1
    ( pr-lg       , padding-right          , $emma-padding-lg ),    // ^1
    ( pr-xl       , padding-right          , $emma-padding-xl ),    // ^1
    ( pb-0        , padding-bottom         , 0 ),
    ( pb-xs       , padding-bottom         , $emma-padding-xs ),    // ^1
    ( pb-sm       , padding-bottom         , $emma-padding-sm ),    // ^1
    ( pb-md       , padding-bottom         , $emma-padding-md ),    // ^1
    ( pb-lg       , padding-bottom         , $emma-padding-lg ),    // ^1
    ( pb-xl       , padding-bottom         , $emma-padding-xl ),    // ^1
    ( pl-0        , padding-left           , 0 ),
    ( pl-xs       , padding-left           , $emma-padding-xs ),    // ^1
    ( pl-sm       , padding-left           , $emma-padding-sm ),    // ^1
    ( pl-md       , padding-left           , $emma-padding-md ),    // ^1
    ( pl-lg       , padding-left           , $emma-padding-lg ),    // ^1
    ( pl-xl       , padding-left           , $emma-padding-xl ),    // ^1
    ( w-a         , width                  , auto ),
    ( w-0         , width                  , 0 ),
    ( w-1p        , width                  , 1% ),
    ( w-10p       , width                  , 10% ),
    ( w-20p       , width                  , 20% ),
    ( w-25p       , width                  , 25% ),
    ( w-30p       , width                  , 30% ),
    ( w-33p       , width                  , 33% ),
    ( w-40p       , width                  , 40% ),
    ( w-50p       , width                  , 50% ),
    ( w-60p       , width                  , 60% ),
    ( w-66p       , width                  , 66% ),
    ( w-70p       , width                  , 70% ),
    ( w-75p       , width                  , 75% ),
    ( w-80p       , width                  , 80% ),
    ( w-90p       , width                  , 90% ),
    ( w-100p      , width                  , 100% ),
    ( h-a         , height                 , auto ),
    ( h-0         , height                 , 0 ),
    ( h-100p      , height                 , 100% ),
    ( maw-n       , max-width              , none ),
    ( maw-100p    , max-width              , 100% ),
    ( mah-n       , max-height             , none ),
    ( mah-100p    , max-height             , 100% ),
    ( miw-0       , min-width              , 0 ),
    ( mih-0       , min-height             , 0 ),
    ( ol-n        , outline                , none ),
    ( olw-tn      , outline-width          , thin ),
    ( olw-md      , outline-width          , medium ),
    ( olw-tc      , outline-width          , thick ),
    ( ols-n       , outline-style          , none ),
    ( ols-dt      , outline-style          , dotted ),
    ( ols-ds      , outline-style          , dashed ),
    ( ols-s       , outline-style          , solid ),
    ( ols-db      , outline-style          , double ),
    ( ols-g       , outline-style          , groove ),
    ( ols-r       , outline-style          , ridge ),
    ( ols-i       , outline-style          , inset ),
    ( ols-o       , outline-style          , outset ),
    ( olc-i       , outline-color          , invert ),
    ( bfv-h       , backface-visibility    , hidden ),
    ( bfv-v       , backface-visibility    , visible ),
    ( bd-n        , border                 , none ),
    ( bd-0        , border                 , 0 ),
    ( bdcl-c      , border-collapse        , collapse ),
    ( bdcl-s      , border-collapse        , separate ),
    ( bdc-t       , border-color           , transparent ),
    ( bdc-cc      , border-color           , currentColor ),        // ^1
    ( bdc-white   , border-color           , $emma-color-white ),   // ^1
    ( bdc-silver  , border-color           , $emma-color-silver ),  // ^1
    ( bdc-gray    , border-color           , $emma-color-gray ),    // ^1
    ( bdc-black   , border-color           , $emma-color-black ),   // ^1
    ( bdc-navy    , border-color           , $emma-color-navy ),    // ^1
    ( bdc-blue    , border-color           , $emma-color-blue ),    // ^1
    ( bdc-aqua    , border-color           , $emma-color-aqua ),    // ^1
    ( bdc-teal    , border-color           , $emma-color-teal ),    // ^1
    ( bdc-olive   , border-color           , $emma-color-olive ),   // ^1
    ( bdc-green   , border-color           , $emma-color-green ),   // ^1
    ( bdc-lime    , border-color           , $emma-color-lime ),    // ^1
    ( bdc-yellow  , border-color           , $emma-color-yellow ),  // ^1
    ( bdc-orange  , border-color           , $emma-color-orange ),  // ^1
    ( bdc-red     , border-color           , $emma-color-red ),     // ^1
    ( bdc-maroon  , border-color           , $emma-color-maroon ),  // ^1
    ( bdc-fuchsia , border-color           , $emma-color-fuchsia ), // ^1
    ( bdc-purple  , border-color           , $emma-color-purple ),  // ^1
    ( bdi-n       , border-image           , none ),
    ( bds-n       , border-style           , none ),
    ( bds-h       , border-style           , hidden ),
    ( bds-dt      , border-style           , dotted ),
    ( bds-ds      , border-style           , dashed ),
    ( bds-s       , border-style           , solid ),
    ( bds-db      , border-style           , double ),
    ( bds-w       , border-style           , wave ),
    ( bds-g       , border-style           , groove ),
    ( bds-r       , border-style           , ridge ),
    ( bds-i       , border-style           , inset ),
    ( bds-o       , border-style           , outset ),
    ( bdw-0       , border-width           , 0 ),
    ( bdw-1       , border-width           , 1px ),
    ( bdw-2       , border-width           , 2px ),
    ( bdw-3       , border-width           , 3px ),
    ( bdw-4       , border-width           , 4px ),
    ( bdw-5       , border-width           , 5px ),
    ( bdw-6       , border-width           , 6px ),
    ( bdtw-0      , border-top-width       , 0 ),
    ( bdtw-1      , border-top-width       , 1px ),
    ( bdtw-2      , border-top-width       , 2px ),
    ( bdtw-3      , border-top-width       , 3px ),
    ( bdtw-4      , border-top-width       , 4px ),
    ( bdtw-5      , border-top-width       , 5px ),
    ( bdtw-6      , border-top-width       , 6px ),
    ( bdrw-0      , border-right-width     , 0 ),
    ( bdrw-1      , border-right-width     , 1px ),
    ( bdrw-2      , border-right-width     , 2px ),
    ( bdrw-3      , border-right-width     , 3px ),
    ( bdrw-4      , border-right-width     , 4px ),
    ( bdrw-5      , border-right-width     , 5px ),
    ( bdrw-6      , border-right-width     , 6px ),
    ( bdbw-0      , border-bottom-width    , 0 ),
    ( bdbw-1      , border-bottom-width    , 1px ),
    ( bdbw-2      , border-bottom-width    , 2px ),
    ( bdbw-3      , border-bottom-width    , 3px ),
    ( bdbw-4      , border-bottom-width    , 4px ),
    ( bdbw-5      , border-bottom-width    , 5px ),
    ( bdbw-6      , border-bottom-width    , 6px ),
    ( bdlw-0      , border-left-width      , 0 ),
    ( bdlw-1      , border-left-width      , 1px ),
    ( bdlw-2      , border-left-width      , 2px ),
    ( bdlw-3      , border-left-width      , 3px ),
    ( bdlw-4      , border-left-width      , 4px ),
    ( bdlw-5      , border-left-width      , 5px ),
    ( bdlw-6      , border-left-width      , 6px ),
    ( bdt-n       , border-top             , none ),
    ( bdt-0       , border-top             , 0 ),
    ( bdtc-t      , border-top-color       , transparent ),
    ( bdtc-cc     , border-top-color       , currentColor ),        // ^1
    ( bdr-n       , border-right           , none ),
    ( bdr-0       , border-right           , 0 ),
    ( bdrc-t      , border-right-color     , transparent ),
    ( bdrc-cc     , border-right-color     , currentColor ),        // ^1
    ( bdb-n       , border-bottom          , none ),
    ( bdb-0       , border-bottom          , 0 ),
    ( bdbc-t      , border-bottom-color    , transparent ),
    ( bdbc-cc     , border-bottom-color    , currentColor ),        // ^1
    ( bdl-n       , border-left            , none ),
    ( bdl-0       , border-left            , 0 ),
    ( bdlc-t      , border-left-color      , transparent ),
    ( bdlc-cc     , border-left-color      , currentColor ),        // ^1
    ( bdrs-0      , border-radius          , 0 ),
    ( bdrs-1      , border-radius          , 1px ),
    ( bdrs-2      , border-radius          , 2px ),
    ( bdrs-3      , border-radius          , 3px ),
    ( bdrs-4      , border-radius          , 4px ),
    ( bdrs-5      , border-radius          , 5px ),
    ( bdrs-6      , border-radius          , 6px ),
    ( bdrs-100p   , border-radius          , 100% ),
    ( bg-n        , background             , none ),
    ( bgc-t       , background-color       , transparent ),
    ( bgc-cc      , background-color       , currentColor ),        // ^1
    ( bgc-white   , background-color       , $emma-color-white ),   // ^1
    ( bgc-silver  , background-color       , $emma-color-silver ),  // ^1
    ( bgc-gray    , background-color       , $emma-color-gray ),    // ^1
    ( bgc-black   , background-color       , $emma-color-black ),   // ^1
    ( bgc-navy    , background-color       , $emma-color-navy ),    // ^1
    ( bgc-blue    , background-color       , $emma-color-blue ),    // ^1
    ( bgc-aqua    , background-color       , $emma-color-aqua ),    // ^1
    ( bgc-teal    , background-color       , $emma-color-teal ),    // ^1
    ( bgc-olive   , background-color       , $emma-color-olive ),   // ^1
    ( bgc-green   , background-color       , $emma-color-green ),   // ^1
    ( bgc-lime    , background-color       , $emma-color-lime ),    // ^1
    ( bgc-yellow  , background-color       , $emma-color-yellow ),  // ^1
    ( bgc-orange  , background-color       , $emma-color-orange ),  // ^1
    ( bgc-red     , background-color       , $emma-color-red ),     // ^1
    ( bgc-maroon  , background-color       , $emma-color-maroon ),  // ^1
    ( bgc-fuchsia , background-color       , $emma-color-fuchsia ), // ^1
    ( bgc-purple  , background-color       , $emma-color-purple ),  // ^1
    ( bgi-n       , background-image       , none ),
    ( bgr-n       , background-repeat      , no-repeat ),
    ( bgr-x       , background-repeat      , repeat-x ),
    ( bgr-y       , background-repeat      , repeat-y ),
    ( bgr-sp      , background-repeat      , space ),
    ( bgr-rd      , background-repeat      , round ),
    ( bga-f       , background-attachment  , fixed ),
    ( bga-s       , background-attachment  , scroll ),
    ( bgp-t       , background-position    , top ),                 // ^1
    ( bgp-r       , background-position    , right ),               // ^1
    ( bgp-b       , background-position    , bottom ),              // ^1
    ( bgp-l       , background-position    , left ),                // ^1
    ( bgp-c       , background-position    , center ),              // ^1
    ( bgsz-a      , background-size        , auto ),
    ( bgsz-ct     , background-size        , contain ),
    ( bgsz-cv     , background-size        , cover ),
    ( c-i         , color                  , inherit ),             // ^1
    ( c-white     , color                  , $emma-color-white ),   // ^1
    ( c-silver    , color                  , $emma-color-silver ),  // ^1
    ( c-gray      , color                  , $emma-color-gray ),    // ^1
    ( c-black     , color                  , $emma-color-black ),   // ^1
    ( c-navy      , color                  , $emma-color-navy ),    // ^1
    ( c-blue      , color                  , $emma-color-blue ),    // ^1
    ( c-aqua      , color                  , $emma-color-aqua ),    // ^1
    ( c-teal      , color                  , $emma-color-teal ),    // ^1
    ( c-olive     , color                  , $emma-color-olive ),   // ^1
    ( c-green     , color                  , $emma-color-green ),   // ^1
    ( c-lime      , color                  , $emma-color-lime ),    // ^1
    ( c-yellow    , color                  , $emma-color-yellow ),  // ^1
    ( c-orange    , color                  , $emma-color-orange ),  // ^1
    ( c-red       , color                  , $emma-color-red ),     // ^1
    ( c-maroon    , color                  , $emma-color-maroon ),  // ^1
    ( c-fuchsia   , color                  , $emma-color-fuchsia ), // ^1
    ( c-purple    , color                  , $emma-color-purple ),  // ^1
    ( tbl-a       , table-layout           , auto ),
    ( tbl-f       , table-layout           , fixed ),
    ( lis-n       , list-style             , none ),
    ( lisp-i      , list-style-position    , inside ),
    ( lisp-o      , list-style-position    , outside ),
    ( list-n      , list-style-type        , none ),
    ( list-d      , list-style-type        , disc ),
    ( list-c      , list-style-type        , circle ),
    ( list-s      , list-style-type        , square ),
    ( list-dc     , list-style-type        , decimal ),
    ( list-dclz   , list-style-type        , decimal-leading-zero ),
    ( list-lr     , list-style-type        , lower-roman ),
    ( list-ur     , list-style-type        , upper-roman ),
    ( lisi-n      , list-style-image       , none ),
    ( va-sup      , vertical-align         , super ),
    ( va-t        , vertical-align         , top ),
    ( va-tt       , vertical-align         , text-top ),
    ( va-m        , vertical-align         , middle ),
    ( va-bl       , vertical-align         , baseline ),
    ( va-b        , vertical-align         , bottom ),
    ( va-tb       , vertical-align         , text-bottom ),
    ( va-sub      , vertical-align         , sub ),
    ( ta-l        , text-align             , left ),
    ( ta-c        , text-align             , center ),
    ( ta-r        , text-align             , right ),
    ( ta-j        , text-align             , justify ),
    ( td-n        , text-decoration        , none ),
    ( td-u        , text-decoration        , underline ),
    ( td-o        , text-decoration        , overline ),
    ( td-l        , text-decoration        , line-through ),
    ( te-n        , text-emphasis          , none ),
    ( te-ac       , text-emphasis          , accent ),
    ( te-dt       , text-emphasis          , dot ),
    ( te-c        , text-emphasis          , circle ),
    ( te-ds       , text-emphasis          , disc ),
    ( te-b        , text-emphasis          , before ),
    ( te-a        , text-emphasis          , after ),
    ( ti-0        , text-indent            , 0 ),
    ( ti--9999    , text-indent            , -9999px ),             // Emmet: ti-
    ( tov-e       , text-overflow          , ellipsis ),
    ( tov-c       , text-overflow          , clip ),
    ( tt-n        , text-transform         , none ),
    ( tt-c        , text-transform         , capitalize ),
    ( tt-u        , text-transform         , uppercase ),
    ( tt-l        , text-transform         , lowercase ),
    ( tsh-n       , text-shadow            , none ),
    ( lh-nm       , line-height            , normal ),
    ( lh-i        , line-height            , inherit ),
    ( lh-0        , line-height            , 0 ),
    ( lh-1        , line-height            , 1 ),
    ( lh-2        , line-height            , 2 ),
    ( lh-3        , line-height            , 3 ),
    ( lh-4        , line-height            , 4 ),
    ( lh-5        , line-height            , 5 ),
    ( lh-6        , line-height            , 6 ),
    ( lh-xs       , line-height            , $emma-line-height-xs ),
    ( lh-sm       , line-height            , $emma-line-height-sm ),
    ( lh-md       , line-height            , $emma-line-height-md ),
    ( lh-lg       , line-height            , $emma-line-height-lg ),
    ( lh-xl       , line-height            , $emma-line-height-xl ),
    ( whs-nm      , white-space            , normal ),              // Emmet: whs:n
    ( whs-p       , white-space            , pre ),
    ( whs-nw      , white-space            , nowrap ),
    ( whs-pw      , white-space            , pre-wrap ),
    ( whs-pl      , white-space            , pre-line ),
    ( wob-nm      , word-break             , normal ),              // Emmet: wob:n
    ( wob-k       , word-break             , keep-all ),
    ( wob-ba      , word-break             , break-all ),
    ( wow-nm      , word-wrap              , normal ),
    ( wow-n       , word-wrap              , none ),
    ( wow-u       , word-wrap              , unrestricted ),
    ( wow-s       , word-wrap              , suppress ),
    ( wow-bw      , word-wrap              , break-word ),
    ( lts-nm      , letter-spacing         , normal ),              // Emmet: lts:n
    ( fw-nm       , font-weight            , normal ),              // Emmet: fw:n
    ( fw-b        , font-weight            , bold ),
    ( fw-br       , font-weight            , bolder ),
    ( fw-l        , font-weight            , 200 ),                 // ^1
    ( fw-lr       , font-weight            , lighter ),
    ( fw-100      , font-weight            , 100 ),
    ( fw-200      , font-weight            , 200 ),
    ( fw-300      , font-weight            , 300 ),
    ( fw-400      , font-weight            , 400 ),
    ( fw-500      , font-weight            , 500 ),
    ( fw-600      , font-weight            , 600 ),
    ( fw-700      , font-weight            , 700 ),
    ( fw-800      , font-weight            , 800 ),
    ( fw-900      , font-weight            , 900 ),
    ( fs-nm       , font-style             , normal ),              // Emmet: fs:n
    ( fs-i        , font-style             , italic ),
    ( fs-o        , font-style             , oblique ),
    ( fv-nm       , font-variant           , normal ),              // Emmet: fv:n
    ( fv-sc       , font-variant           , small-caps ),
    ( fz-xs       , font-size              , $emma-font-size-xs ),  // ^1
    ( fz-sm       , font-size              , $emma-font-size-sm ),  // ^1
    ( fz-md       , font-size              , $emma-font-size-md ),  // ^1
    ( fz-lg       , font-size              , $emma-font-size-lg ),  // ^1
    ( fz-xl       , font-size              , $emma-font-size-xl ),  // ^1
    ( fz-sr       , font-size              , smaller ),             // ^1
    ( fz-lr       , font-size              , larger ),              // ^1
    ( fz-10       , font-size              , 10px ),
    ( fz-11       , font-size              , 11px ),
    ( fz-12       , font-size              , 12px ),
    ( fz-13       , font-size              , 13px ),
    ( fz-14       , font-size              , 14px ),
    ( fz-15       , font-size              , 15px ),
    ( fz-16       , font-size              , 16px ),
    ( fz-17       , font-size              , 17px ),
    ( fz-18       , font-size              , 18px ),
    ( fz-19       , font-size              , 19px ),
    ( fz-20       , font-size              , 20px ),
    ( fz-h1       , font-size              , $emma-font-size-h1 ),  // ^1
    ( fz-h2       , font-size              , $emma-font-size-h2 ),  // ^1
    ( fz-h3       , font-size              , $emma-font-size-h3 ),  // ^1
    ( fz-h4       , font-size              , $emma-font-size-h4 ),  // ^1
    ( fz-h5       , font-size              , $emma-font-size-h5 ),  // ^1
    ( fz-h6       , font-size              , $emma-font-size-h6 ),  // ^1
    ( ff-s        , font-family            , serif ),
    ( ff-ss       , font-family            , sans-serif ),
    ( ff-c        , font-family            , cursive ),
    ( ff-f        , font-family            , fantasy ),
    ( ff-m        , font-family            , monospace ),
    ( ff-a        , font-family            , 'Arial, "Helvetica Neue", Helvetica, sans-serif' ),
    ( ff-t        , font-family            , '"Times New Roman", Times, Baskerville, Georgia, serif' ),
    ( ff-v        , font-family            , 'Verdana, Geneva, sans-serif' ),
    ( ff-l        , font-family            , '"Lucida Grande", "Lucida Sans Unicode", Verdana, Arial, "Helvetica Neue", Helvetica, sans-serif' ),
    ( op-0        , opacity                , 0 ),
    ( op-0_1      , opacity                , 0.1 ),                 // ^1
    ( op-0_2      , opacity                , 0.2 ),                 // ^1
    ( op-0_3      , opacity                , 0.3 ),                 // ^1
    ( op-0_4      , opacity                , 0.4 ),                 // ^1
    ( op-0_5      , opacity                , 0.5 ),                 // ^1
    ( op-0_6      , opacity                , 0.6 ),                 // ^1
    ( op-0_7      , opacity                , 0.7 ),                 // ^1
    ( op-0_8      , opacity                , 0.8 ),                 // ^1
    ( op-0_9      , opacity                , 0.9 ),                 // ^1
    ( op-1        , opacity                , 1 ),
    ( rsz-n       , resize                 , none ),
    ( rsz-b       , resize                 , both ),
    ( rsz-h       , resize                 , horizontal ),
    ( rsz-v       , resize                 , vertical ),
    ( cur-a       , cursor                 , auto ),
    ( cur-d       , cursor                 , default ),
    ( cur-c       , cursor                 , crosshair ),
    ( cur-ha      , cursor                 , hand ),
    ( cur-he      , cursor                 , help ),
    ( cur-m       , cursor                 , move ),
    ( cur-p       , cursor                 , pointer ),
    ( cur-t       , cursor                 , text ),
    ( fxd-r       , flex-direction         , row ),
    ( fxd-rr      , flex-direction         , row-reverse ),
    ( fxd-c       , flex-direction         , column ),
    ( fxd-cr      , flex-direction         , column-reverse ),
    ( fxw-n       , flex-wrap              , nowrap ),
    ( fxw-w       , flex-wrap              , wrap ),
    ( fxw-wr      , flex-wrap              , wrap-reverse ),
    ( jc-fs       , justify-content        , flex-start ),
    ( jc-fe       , justify-content        , flex-end ),
    ( jc-c        , justify-content        , center ),
    ( jc-sb       , justify-content        , space-between ),
    ( jc-sa       , justify-content        , space-around ),
    ( ai-fs       , align-items            , flex-start ),
    ( ai-fe       , align-items            , flex-end ),
    ( ai-c        , align-items            , center ),
    ( ai-b        , align-items            , baseline ),
    ( ai-s        , align-items            , stretch ),
    ( ac-fs       , align-content          , flex-start ),
    ( ac-fe       , align-content          , flex-end ),
    ( ac-c        , align-content          , center ),
    ( ac-sb       , align-content          , space-between ),
    ( ac-sa       , align-content          , space-around ),
    ( ac-s        , align-content          , stretch ),
    ( ord--1      , order                  , -1 ),
    ( ord-0       , order                  , 0 ),
    ( ord-1       , order                  , 1 ),
    ( ord-2       , order                  , 2 ),
    ( ord-3       , order                  , 3 ),
    ( ord-4       , order                  , 4 ),
    ( ord-5       , order                  , 5 ),
    ( ord-6       , order                  , 6 ),
    ( ord-9999    , order                  , 9999 ),
    ( fx-n        , flex                   , none ),
    ( fx-1_1_a    , flex                   , "1 1 auto" ),          // ^1
    ( fx-1_0_a    , flex                   , "1 0 auto" ),          // ^1
    ( fx-1_1_1    , flex                   , "1 1 1px" ),           // ^1
    ( as-a        , align-self             , auto ),
    ( as-fs       , align-self             , flex-start ),
    ( as-fe       , align-self             , flex-end ),
    ( as-c        , align-self             , center ),
    ( as-b        , align-self             , baseline ),
    ( as-s        , align-self             , stretch ),
    ( wfsm-a      , -webkit-font-smoothing , antialiased ),
    ( wfsm-sa     , -webkit-font-smoothing , subpixel-antialiased ),
    ( wfsm-n      , -webkit-font-smoothing , none ),
    ( obf-f       , object-fit             , fill ),                // ^1
    ( obf-ct      , object-fit             , contain ),             // ^1
    ( obf-cv      , object-fit             , cover ),               // ^1
    ( obf-n       , object-fit             , none ),                // ^1
    ( obf-sd      , object-fit             , scale-down ),          // ^1
`
