package area

type Area struct {
    KixcodeBegin        string `json:"kixcodeBegin"`
    KixcodeEndInclusive string `json:"kixcodeEndInclusive"`
}

type AreaStore struct {
    Areas []Area
};

func (areaStore *AreaStore) Add(newArea Area) {
    for _, area := range areaStore.Areas {
        if area.KixcodeBegin == newArea.KixcodeBegin && area.KixcodeEndInclusive == newArea.KixcodeEndInclusive {
            return
        }
    }
    areaStore.Areas = append(areaStore.Areas, newArea);
}

func (areaStore *AreaStore) Remove(kixcodeBegin, kixcodeEnd string) {
    for i, area := range areaStore.Areas {
        if area.KixcodeBegin == kixcodeBegin && area.KixcodeEndInclusive == kixcodeEnd {
            areaStore.Areas = append(areaStore.Areas[:i], areaStore.Areas[i + 1:]...);
            return
        }
    }
}

func (areaStore *AreaStore) Get(kixcodeBegin, kixcodeEnd string) (Area, bool) {
    for _, area := range areaStore.Areas {
        if area.KixcodeBegin == kixcodeBegin && area.KixcodeEndInclusive == kixcodeEnd {
            return area, true
        }
    }
    return Area{}, false
}

func (areaStore *AreaStore) Excludes(kixcode string) bool {
    for _, area := range areaStore.Areas {
        if area.KixcodeBegin <= kixcode && kixcode <= area.KixcodeEndInclusive {
            return true
        }
    }
    return false
}

