package models

type ScoringKey struct{
    FullNum string `csv:"Full#"`
    ShortNum string `csv:"Short#"`
    Sign string `csv:"Sign"`
    Key string `csv:"Key"`
    Facet string `csv:"Facet"`
    Item string `csv:"Item"`
}
