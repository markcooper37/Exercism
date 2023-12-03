class RnaTranscription {
    
    def static transcriptions = ['G': 'C', 'C': 'G', 'T': 'A', 'A': 'U']

    static String toRna(String strand) {
        return strand.collect{transcriptions[it]}.join()
    }

}
