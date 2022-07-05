"""
ChineseParams: the params for the NLP pipeline.
"""


class ChineseParams:

    """
    feature_map, stop word
    """

    def __init__(self, feature_map: dict):

        # feature map
        self.feature_map = {
            "word_segmentation": feature_map.get("word_segmentation", True),
            "stop_word_removal": feature_map.get("stop_word_removal", False),
        }

        # stop word
        self.stop_word_set = set()

    def set_feature_map(self, key: str, value: bool) -> bool:
        """update feature map; will check feature map key; return success boolean"""
        if key not in self.feature_map:
            return False
        self.feature_map[key] = value
        return True

    def add_stop_words(self, stop_words: set) -> bool:
        """merge stop words from input set"""
        self.stop_word_set |= stop_words
        return True

    def remove_stop_words(self, stop_words: set) -> bool:
        """remove stop words from input set"""
        self.stop_word_set -= stop_words
        return True

    def clear_stop_words(self) -> bool:
        """empty stop word set"""
        self.stop_word_set = set()
        return True
